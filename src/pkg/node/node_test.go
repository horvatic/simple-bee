package node

import ( 
	"testing"
	"reflect"
)

func TestGets(t *testing.T) {
	sampleTags := []string{"arm", "x64"}
	sampleNode := NewNode("1.1.1.0", "test", sampleTags)

    if sampleNode.GetIp() != "1.1.1.0" {
        t.Errorf("GetIp() = %v; 1.1.1.0", sampleNode.GetIp())
    }
	if sampleNode.GetName() != "test" {
        t.Errorf("GetName() = %v; test", sampleNode.GetName())
    }
	if !reflect.DeepEqual(sampleNode.GetTags(), []string{"arm", "x64"}) {
        t.Errorf("GetTags() = %v; [arm x64]", sampleNode.GetTags())
    }
}