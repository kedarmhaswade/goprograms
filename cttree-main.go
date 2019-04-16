/*
	This standalone program is intended to aid the visualization of the contact type and contact type group hierarchies.

 	It builds on top of two endpoints provided by the contact service's IDL:
 		1) list<ContactTypeGroup> getAllContactTypeGroups(), and
 		2) list<ContactTypeNode> getAllContactTypeNodesV2() (you can include the deleted ones in the request JSON)

	As such, a caller should first obtain the output of these two endpoints (e.g. by running yab)
	and feed that to this program as command-line arguments to get the JSON files that help visualize
	the contact type and contact type group hierarchies.

	e.g. cttree -ctgin ctgin.json -ctnin ctnin.json < -ctgout ctgout.json > < --ctnout ctnout.json >

 	will accept the group and node information in the form of ctgs.json and ctn.json and produce groups.json and
	nodes.json (in the current working directory by default)
 	--help would produce this help message and exit.
    -ctgout and -ctnout are optional.
 	The generated files (e.g. ctgout.json and ctnout.json) can be viewed/visualized using a tool of your choice.

*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// We treat ROOT_ID as the ID of the "implicit" root of contact type hierarchy
const ROOT_ID = "00000000-0000-0000-0000-000000000000"
const ROOT_NAME = "root"

// Structures to consume the output of yab commands that is a JSON object of the form:
// {"body": {"result":[]}, "ok": <boolean>, "trace": <string>}
// Strictly speaking, this program should be given a valid response (e.g. "ok": true)
// None of these structures are exported from this package.

// ctn yab
var ctnYab = "yab -s contact_service -m ContactService::getAllContactTypeNodesV2"
// ctnYabResponse defines the output of the yab command that gets all contact type nodes
type ctnYabResponse struct {
	Body  ctnYabResponseBody `json:"body"`
	Ok    bool               `json:"ok"`
	Trace string             `json:"trace"`
}

// ctnYabResponseBody defines the "body" of the yab command that gets all contact type nodes
type ctnYabResponseBody struct {
	Result []inNode `json:"result"`
}

// ctg yab
var ctgYab = "yab -s contact_service -m ContactService::getAllContactTypeGroups"
// ctgYabResponse defines the output of the yab command that gets all contact type groups
type ctgYabResponse struct {
	Body  ctgYabResponseBody `json:"body"`
	Ok    bool               `json:"ok"`
	Trace string             `json:"trace"`
}
// ctgYabResponseBody defines the "body" of the yab command that gets all contact type groups
type ctgYabResponseBody struct {
	Result []*groupNode `json:"result"`
}

// Structures that deal with contact type tree and contact type groups

// inNode defines the contact type tree node in the input to this program; it is not recursive
type inNode struct {
	ID        string
	Name      string   `json:"name"`
	Children  []string `json:"childrenIDs"`
	IsDeleted bool     `json:"isDeleted"`
	ParentID  string   `json:"parentID"`
}

//groupNode defines the contact type group node; the structure is self-explanatory
type groupNode struct {
	ID                 string
	Name               string   `json:"name"`
	ContactTypeNodeIDs []string `json:",omitempty"`
	Members            []*treeNode
}

//treeNode defines the *recursive* data structure that determines the contact type and contact type group hierarchy
type treeNode struct {
	ID        string
	Name      string
	IsDeleted bool
	Children  []*treeNode `json:",omitempty"`
}

//ROOT_NODE, an instance of treeNode, is the root of contact type node hierarchy
var ROOT_NODE = treeNode{
	ID:        ROOT_ID,
	Name:      ROOT_NAME,
	IsDeleted: false,
}

// Flag variables (package scope)
// file name of response of getAllContactTypeGroups()
var ctgInVar string
// file name of response of getAllContactTypeNodesV2()
var ctnInVar string
// file name of the contact type groups file produced by this program
var ctgOutVar string
// file name of the contact type nodes file produced by this program
var ctnOutVar string

func init() {
	const CTGIN = "ctgin"
	const CTNIN = "ctnin"
	const CTGOUT = "ctgout"
	const CTNOUT = "ctnout"
	flag.StringVar(&ctgInVar, CTGIN, "", "file name of response of getAllContactTypeGroups()")
	flag.StringVar(&ctnInVar, CTNIN, "", "file name of response of getAllContactTypeNodesV2()")
	flag.StringVar(&ctgOutVar, CTGOUT, CTGOUT+".json", "file name of the contact type groups file produced by this program")
	flag.StringVar(&ctnOutVar, CTNOUT, CTNOUT+".json", "file name of the contact type nodes file produced by this program")
	flag.Parse()
	cmdSyntaxOk := true
	if len(ctgInVar) == 0 {
		fmt.Fprintf(os.Stderr, "Error: -%s is required\n", CTGIN)
		cmdSyntaxOk = false
	}
	if len(ctnInVar) == 0 {
		fmt.Fprintf(os.Stderr, "Error: -%s is required\n", CTNIN)
		cmdSyntaxOk = false
	}
	if !cmdSyntaxOk {
		flag.Usage()
		os.Exit(1)
	}
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {

	var data []byte
	data, e := ioutil.ReadFile(ctnInVar)
	if e != nil {
		log.Fatalf("error: %v reading file %s", e, ctnInVar)
	}
	// ctn first
	var ctnResponse = ctnYabResponse{}
	e = json.Unmarshal(data, &ctnResponse)
	if e != nil {
		log.Fatalf("error %v while unmarshaling into contact type node structure\nRun %s first", e, ctnYab)
	}
	var inNodes = ctnResponse.Body.Result
	log.Printf("number of nodes: %v\n", len(inNodes))
	inMap := make(map[string]*inNode, len(inNodes))
	outMap := make(map[string]*treeNode, len(inNodes))
	rootNode := buildTree(inNodes, inMap, outMap)
	tree, e := json.Marshal(rootNode)
	if e != nil {
		log.Fatalf("error [%v] while marshaling contact type tree to JSON\n", e)
	}
	e = ioutil.WriteFile(ctnOutVar, tree, 0644) // write with perm of regular file
	if e != nil {
		log.Fatalf("error [%v] while writing to file %s\n", e, ctnOutVar)
	}
	// ctg next
	var ctgResponse = ctgYabResponse{}
	data, e = ioutil.ReadFile(ctgInVar)
	if e != nil {
		log.Fatalf("error: [%v] while reading file %s", e, ctgInVar)
	}
	e = json.Unmarshal(data, &ctgResponse)
	if e != nil {
		log.Fatalf("error [%v] while unmarshaling into contact type group structure\nRun %s first", e, ctgYab)
	}
	var groupNodes = ctgResponse.Body.Result
	for _, gNode := range groupNodes {
		for _, ctNodeID := range gNode.ContactTypeNodeIDs {
			outNode, ok := outMap[ctNodeID]
			if !ok {
				_, _ = fmt.Fprintf(os.Stderr, "warning: outNode id [%v] not found for groupNode: [%v]\n", outNode.ID, gNode.ID)
				continue
			}
			gNode.Members = append(gNode.Members, outNode)
		}
		gNode.ContactTypeNodeIDs = nil // release, let the gc take care of it if needed
	}
	tree, e = json.Marshal(groupNodes)
	if e != nil {
		log.Fatalf("error [%v] while marshaling the contact type group structure\n", e)
	}
	e = ioutil.WriteFile(ctgOutVar, tree, 0644) // write with perm of regular file
	if e != nil {
		log.Fatalf("error [%v] while writing to file %s\n", e, ctgOutVar)
	}
}

func buildTree(nodes []inNode, inMap map[string]*inNode, outMap map[string]*treeNode) *treeNode {
	fillMaps(nodes, inMap, outMap)
	outMap[ROOT_ID] = &ROOT_NODE // add root explicitly
	for _, node := range nodes {
		me, ok := outMap[node.ID]
		if !ok {
			_ = fmt.Errorf("*odd*, node [id: %v], [name: %v] not found\n", node.ID, node.Name)
			continue // abort?
		}
		p, ok := outMap[node.ParentID]
		if !ok {
			_, _ = fmt.Fprintf(os.Stderr, "warning: parent [%v] for node [%v] not found\n", node.ParentID, node.ID)
			continue
		}
		p.Children = append(p.Children, me)
		//fmt.Printf("handled %v\n", me.ID)
	}
	return &ROOT_NODE
}

func fillMaps(nodes []inNode, inMap map[string]*inNode, outMap map[string]*treeNode) {
	for _, inNode := range nodes {
		inMap[inNode.ID] = &inNode
		outMap[inNode.ID] = newEmptyTreeNode(inNode)
	}
}

func newEmptyTreeNode(in inNode) *treeNode {
	outNode := treeNode{
		Name:      in.Name,
		ID:        in.ID,
		IsDeleted: in.IsDeleted,
		Children:  nil,
	}
	return &outNode
}

//UnMarshalContactTreeNodes reads the given JSON of contact type nodes and deserializes it into the given slice
func UnMarshalContactTreeNodes(reader io.Reader, nodes []inNode) {

}
