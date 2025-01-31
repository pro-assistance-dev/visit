package routing

import (
	"portal/handlers/auth"
	"portal/handlers/banners"
	"portal/handlers/customsections"
	"portal/handlers/event"
	"portal/handlers/eventdays"
	"portal/handlers/eventmessages"
	"portal/handlers/experiences"
	"portal/handlers/m2m"
	"portal/handlers/perfoms"
	"portal/handlers/place"
	"portal/handlers/schedules"
	"portal/handlers/sessions"
	"portal/handlers/speakers"
	"portal/handlers/sponsors"
	"portal/handlers/users"
	"portal/handlers/userseventssactivities"
	authRouter "portal/routing/auth"
	bannersRouter "portal/routing/banners"
	customsectionsRouter "portal/routing/customsections"
	eventdaysRouter "portal/routing/eventdays"
	eventMessagesRouter "portal/routing/eventmessages"
	eventsRouter "portal/routing/events"
	experiencesRouter "portal/routing/experiences"
	m2mRouter "portal/routing/m2m"
	perfomsRouter "portal/routing/perfoms"
	placeRouter "portal/routing/place"
	schedulesRouter "portal/routing/schedules"
	sessionsRouter "portal/routing/sessions"
	speakersRouter "portal/routing/speakers"
	sponsorsRouter "portal/routing/sponsors"
	usersRouter "portal/routing/users"
	usersEventsActivitiesRouter "portal/routing/userseventsactivities"

	"github.com/gin-gonic/gin"
	helperPack "github.com/pro-assistance-dev/sprob/helper"
	"github.com/pro-assistance-dev/sprob/middleware"
	baseRouter "github.com/pro-assistance-dev/sprob/routing"
)

func Init(r *gin.Engine, helper *helperPack.Helper) {
	m := middleware.CreateMiddleware(helper)
	api, apiNoToken := baseRouter.Init(r, helper)
	api.Use(m.InjectFTSP())
	ws := r.Group("/ws")

	auth.Init(helper)
	authRouter.Init(apiNoToken.Group("/auth"), auth.H)

	event.Init(helper)
	eventsRouter.Init(api.Group("/events"), event.H, ws.Group("/events"))

	sponsors.Init(helper)
	sponsorsRouter.Init(api.Group("/sponsors"), sponsors.H)

	banners.Init(helper)
	bannersRouter.Init(api.Group("/banners"), banners.H)

	// humans.Init(helper)
	// humansRouter.Init(api.Group("/humans"), humans.H)

	schedules.Init(helper)
	schedulesRouter.Init(api.Group("/schedules"), schedules.H)

	m2m.Init(helper)
	m2mRouter.Init(api.Group("/m2m"), m2m.H)

	// menus.Init(helper)
	// menusRouter.Init(api.Group("/menus"), menus.H)

	perfoms.Init(helper)
	perfomsRouter.Init(api.Group("/perfoms"), perfoms.H)

	speakers.Init(helper)
	speakersRouter.Init(api.Group("/speakers"), speakers.H)

	eventdays.Init(helper)
	eventdaysRouter.Init(api.Group("/event-days"), eventdays.H)

	place.Init(helper)
	placeRouter.Init(api.Group("/places"), place.H)

	experiences.Init(helper)
	experiencesRouter.Init(api.Group("/experiences"), experiences.H)

	customsections.Init(helper)
	customsectionsRouter.Init(api.Group("/custom-sections"), customsections.H)

	eventmessages.Init(helper)
	eventMessagesRouter.Init(api.Group("/event-messages"), eventmessages.H)

	usersEventsActivitiesRouter.Init(api.Group("/users-events-activities"), userseventssactivities.CreateHandler(helper))

	users.Init(helper)
	usersRouter.Init(api.Group("/users"), users.H)

	//
	// contacts.Init(helper)
	// contactsRouter.Init(api.Group("/contacts"), contacts.H)
	//
	// emails.Init(helper)
	// emailsRouter.Init(api.Group("/emails"), emails.H)

	// phones.Init(helper)
	// phonesRouter.Init(api.Group("/phones"), phones.H)

	sessions.Init(helper)
	sessionsRouter.Init(api.Group("/sessions"), sessions.H)
}
