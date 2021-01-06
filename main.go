package main

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
	"github.com/pseudomuto/protokit/utils"
	_ "google.golang.org/genproto/googleapis/api/annotations" // Support (google.api.http) option (from google/api/annotations.proto).
	"log"
)
/**
protoc插件开发说明

protoc原理：
	第一步，protoc对.proto文件进行解析，生成FileDescriptorSet信息
	第二步，protoc通过os.stdin把FileDescriptorSet信息传递给插件
	第三步，插件通过os.stdin获取到FileDescriptorSet信息，生成plugin_go.CodeGeneratorResponse_File，并通过os.stdout传回给protoc
	第四步，protoc根据 --xxplugin-out 参数配置的目录生成插件指定的文件

protoc插件命令说明：
	例如：protoc --bdw_out=. --go_out=. --go-grpc_out=. api.proto
			protoc在path环境变量中寻找对应的3个插件 protoc-gen-bdw.exe，protoc-gen-go.exe，protoc-gen-go-grpc.exe
			并将FileDescriptorSet信息传给插件，获得插件的输出，从而生成输出文件

protoc插件golang开发：
	开发插件有一个标准的实现模式，参见olddemo文件夹
	标准开发方式比较繁琐，本次开发使用github.com/pseudomuto/protokit包，
	只需要实现func (p *plugin) Generate(in *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error)函数即可
*/
func main() {
	// ****** debug code *******************************
	// protoc正常是通过stdin,stdout传输信息的，不方便调试
	// 我们可以通过如下命令将.proto文件的FileDescriptorSet信息输出到.pb文件中
	// protoc --descriptor_set_out=api.pb --include_imports --include_source_info api.proto
	fds, err := utils.LoadDescriptorSet( "C:\\code\\bdwsms\\api\\fileset.pb")
	if err!=nil{
		println(err.Error())
	}
	req := utils.CreateGenRequest(fds, "api.proto")
	data, err := proto.Marshal(req)
	if err!=nil{
		println(err.Error())
	}

	in := bytes.NewBuffer(data)
	out := new(bytes.Buffer)


	if err!=nil{
		println(err.Error())
	}
	if err := protokit.RunPluginWithIO(new(plugin),in, out); err != nil {
		log.Fatal(err)
	}
	// ****** product code *********************
	// 最终程序使用下面的代码，把debug code的部分全部注释掉
	//if err := protokit.RunPluginWithIO(new(plugin),os.Stdin, os.Stdout); err != nil {
	//	log.Fatal(err)
	//}
}

// 实现插件
type plugin struct{}

func (p *plugin) Generate(in *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	descriptors := protokit.ParseCodeGenRequest(in)

	resp := new(plugin_go.CodeGeneratorResponse)

	for _, d := range descriptors {
		// TODO: YOUR WORK HERE
		fileName := d.GetName()+"zjl" // generate a file name based on d.GetName()
		content := "zjlcontent"// generate content for the output file

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(fileName),
		Content: proto.String(content),
		})
	}

	return resp, nil
}