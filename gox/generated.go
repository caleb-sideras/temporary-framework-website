
// Code generated by gox; DO NOT EDIT.
package gox
import (
	"github.com/caleb-sideras/gox2/src/app/projects/caas"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/processor"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/segmentation"
	"github.com/caleb-sideras/gox2/src/app/_test_"
	"github.com/caleb-sideras/gox2/src/app"
	"github.com/caleb-sideras/gox2/src/app/projects/tweetailyze"
	"github.com/caleb-sideras/gox2/src/app/projects/musicgpt"
	"github.com/caleb-sideras/gox2/src/app/experience"
	"github.com/caleb-sideras/gox2/src/app/projects/peoplepedia"
	"github.com/caleb-sideras/gox2/src/app/home_"
	"github.com/caleb-sideras/gox2/src/app/projects/goxFramework"
	"github.com/caleb-sideras/gox2/src/app/blog"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/cluster"
	"github.com/caleb-sideras/gox2/src/app/blog/temporarystandard"
	"github.com/caleb-sideras/gox2/src/app/projects"
)

var IndexList = map[string]IndexDefaultFunc{
	"/projects/caas/cluster" : app.Index,
	"/" : app.Index,
	"/projects/caas" : app.Index,
	"/projects/caas/segmentation" : app.Index,
	"/projects/tweetailyze" : app.Index,
	"/blog" : app.Index,
	"/projects/caas/processor" : app.Index,
	"/projects/musicgpt" : app.Index,
	"/blog/temporarystandard" : app.Index,
	"/experience" : app.Index,
	"/projects" : app.Index,
	"/projects/peoplepedia" : app.Index,
	"/projects/goxFramework" : app.Index,
	"/{test}" : app.Index,
}

var PageRenderList = []RenderDefault{
	{"/blog", blog.Page_},
	{"/experience", experience.Page_},
}

var RouteRenderList = []RenderDefault{
	{"/example", home.Example_},
}

var PageHandleList = []Handler{
	{"/projects/peoplepedia", peoplepedia.Page, ResReqHandler},
	{"/", home.Page, ResReqHandler},
	{"/projects/caas", caas.Page, ResReqHandler},
	{"/projects/caas/processor", caas_processor.Page, ResReqHandler},
	{"/projects/goxFramework", gox.Page, ResReqHandler},
	{"/projects/caas/segmentation", caas_segmentation.Page, ResReqHandler},
	{"/projects/tweetailyze", tweetailyze.Page, ResReqHandler},
	{"/projects/caas/cluster", caas_cluster.Page, ResReqHandler},
	{"/projects/musicgpt", musicgpt.Page, ResReqHandler},
	{"/{test}", test.Page, ResReqHandler},
	{"/blog/temporarystandard", temporarystandard.Page, ResReqHandler},
	{"/projects", projects.Page, ResReqHandler},
}

var RouteHandleList = []Handler{
	{"/projects/caas/readme", caas.Readme, DefaultHandler},
	{"/projects/caas/videos", caas.Videos, DefaultHandler},
	{"/projects/goxFramework/readme", gox.Readme, DefaultHandler},
	{"/projects/goxFramework/videos", gox.Videos, DefaultHandler},
	{"/projects/musicgpt/readme", musicgpt.Readme, DefaultHandler},
	{"/projects/musicgpt/videos", musicgpt.Videos, DefaultHandler},
}
