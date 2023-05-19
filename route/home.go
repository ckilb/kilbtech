package route

import (
	"ckilb/kilbtech/dto"
	"ckilb/kilbtech/tpl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Home struct {
}

func (r *Home) Path() string {
	return "/"
}

func (r *Home) Method() string {
	return http.MethodGet
}

func (r *Home) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := tpl.TemplateData{
			Projects: []dto.Project{
				{
					LogoFileName: "marel.svg",
					Title:        "Marel",
					Position:     "Technical Lead, Customer Portal",
					Url:          "https://portal.marel.com",
					Description:  "As a <strong>Technical Lead</strong>, I was responsible for creating a new <strong>customer portal</strong> platform from scratch. This new platform enables customers of one of the largest food processing companies worldwide to place, view, and manage their orders online. The successful implementation of this project was a significant step towards digitalization for Marel.",
					LogoWidth:    1730,
					LogoHeight:   410,
				},
				{
					LogoFileName: "aboutyou.svg",
					Title:        "Legal You",
					Position:     "Product Lead, 3rd Party Integration APIs",
					Url:          "https://aboutyou.tech",
					Description:  "As the <strong>Product Lead</strong> for the ABOUT YOU <strong>Developer Center</strong>, I was responsible for enabling third-party apps to connect to the shop ecosystem. I worked closely with a team of six experienced software developers to coordinate the development of the platform, interfaces, and software development kits.",
					LogoWidth:    1024,
					LogoHeight:   214,
				},
				{
					LogoFileName: "dept.svg",
					Title:        "DEPT",
					Position:     "Lead Developer, B2B e-commerce",
					Url:          "https://www.deptagency.com",
					Description:  "As a Lead Developer, I was asked to support DEPT with my expertise in several enterprise-level projects in an advisory and hands-on manner. In close exchange with the customers, I collected, specified and implemented requirements.\n\nThe complexity of the projects often made it necessary to break down extensive business requirements and problems into smaller components - in order to finally communicate them to the team of experienced Spryker software developers in an understandable way",
					LogoWidth:    730,
					LogoHeight:   200,
				},
				{
					LogoFileName: "tomtailor.svg",
					Title:        "Tom Tailor",
					Position:     "Senior Developer & Consultant, B2C e-commerce platform",
					Url:          "https://www.tom-tailor.de",
					Description:  "As a <strong>senior developer and consultant</strong>, I supported the expansion of the Tom Tailor online shop in close cooperation with my colleagues.\nIn addition to my work as a software developer in the implementation, I also created various proofs of concepts, some of which I initiated proactively.\n\n(Junior) developers I was available in an advisory capacity to support them in the implementation of the requirements and further training of their knowledge, especially regarding S.O.L.I.D. coding principles.",
					LogoWidth:    601,
					LogoHeight:   59,
				},
				{
					LogoFileName: "kilb.svg",
					Title:        "Product Designer",
					Position:     "Owner",
					Url:          "https://store.shopware.com/en/kilb647502658205/kilb-product-designer-product-designer-for-shirts-cups-posters-cards...-plugin-version.html",
					Description:  "Along with working for my clients, I continuously strive to enhance my skills and knowledge by working on my own projects. One noteworthy example is the <strong>Product Designer extension</strong> I created for <strong>Shopware</strong> and <strong>WooCommerce</strong>. This extension enables medium-sized and enterprise-level businesses to offer their customers personalized designs for their products.",
					LogoWidth:    1544,
					LogoHeight:   553,
				},
			},
		}

		c.HTML(http.StatusOK, "home", data)
	}
}

func NewHome() Route {
	return &Home{}
}
