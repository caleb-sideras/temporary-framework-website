
// Code generated by gox; DO NOT EDIT.
package gox
import (
	"github.com/caleb-sideras/gox2/src/app/projects/musicgpt"
	"github.com/caleb-sideras/gox2/src/app/projects/peoplepedia"
	"github.com/caleb-sideras/gox2/src/app"
	"github.com/caleb-sideras/gox2/src/app/blog"
	"github.com/caleb-sideras/gox2/src/app/projects"
	"github.com/caleb-sideras/gox2/src/app/projects/goxFramework"
	"github.com/caleb-sideras/gox2/src/app/_test_/rn"
	"github.com/caleb-sideras/gox2/src/app/projects/tweetailyze"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/segmentation"
	"github.com/caleb-sideras/gox2/src/app/blog/temporarystandard"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/processor"
	"github.com/caleb-sideras/gox2/src/app/home_"
	"github.com/caleb-sideras/gox2/src/app/experience"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/cluster"
	"github.com/caleb-sideras/gox2/src/app/_test_"
	"github.com/caleb-sideras/gox2/src/app/projects/caas"
)

var IndexList = map[string]IndexDefaultFunc{
	"/{test}/rn" : app.Index,
	"/projects/caas/processor" : app.Index,
	"/{test}" : app.Index,
	"/projects/musicgpt" : app.Index,
	"/blog/temporarystandard" : app.Index,
	"/projects/tweetailyze" : app.Index,
	"/projects/caas/segmentation" : app.Index,
	"/experience" : app.Index,
	"/blog" : app.Index,
	"/projects/caas" : app.Index,
	"/projects/gox-framework" : app.Index,
	"/projects/caas/cluster" : app.Index,
	"/projects" : app.Index,
	"/projects/peoplepedia" : app.Index,
	"/" : app.Index,
}

var PageRenderList = []RenderDefault{
	{"/experience", experience.Page_},
	{"/blog", blog.Page_},
}

var RouteRenderList = []RenderDefault{
	{"/{test}/example", test.Example_},
	{"/example", home.Example_},
}

var PageHandleList = []Handler{
	{"/{test}/rn", rn.Page, DefaultHandler},
	{"/projects/caas/cluster", caas_cluster.Page, ResReqHandler},
	{"/projects", projects.Page, ResReqHandler},
	{"/projects/caas", caas.Page, ResReqHandler},
	{"/projects/gox-framework", gox.Page, ResReqHandler},
	{"/projects/musicgpt", musicgpt.Page, ResReqHandler},
	{"/projects/peoplepedia", peoplepedia.Page, ResReqHandler},
	{"/blog/temporarystandard", temporarystandard.Page, ResReqHandler},
	{"/projects/caas/processor", caas_processor.Page, ResReqHandler},
	{"/projects/tweetailyze", tweetailyze.Page, ResReqHandler},
	{"/", home.Page, ResReqHandler},
	{"/projects/caas/segmentation", caas_segmentation.Page, ResReqHandler},
}

var RouteHandleList = []Handler{
	{"/{test}/example2", test.Example2, DefaultHandler},
	{"/projects/caas/readme", caas.Readme, DefaultHandler},
	{"/projects/caas/videos", caas.Videos, DefaultHandler},
	{"/projects/gox-framework/readme", gox.Readme, DefaultHandler},
	{"/projects/gox-framework/videos", gox.Videos, DefaultHandler},
	{"/projects/musicgpt/readme", musicgpt.Readme, DefaultHandler},
	{"/projects/musicgpt/videos", musicgpt.Videos, DefaultHandler},
}
