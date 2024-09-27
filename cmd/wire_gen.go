// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package answercmd

import (
	"github.com/apache/incubator-answer/internal/base/conf"
	"github.com/apache/incubator-answer/internal/base/cron"
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/middleware"
	"github.com/apache/incubator-answer/internal/base/server"
	"github.com/apache/incubator-answer/internal/base/translator"
	"github.com/apache/incubator-answer/internal/controller"
	"github.com/apache/incubator-answer/internal/controller/template_render"
	"github.com/apache/incubator-answer/internal/controller_admin"
	"github.com/apache/incubator-answer/internal/repo/activity"
	"github.com/apache/incubator-answer/internal/repo/activity_common"
	"github.com/apache/incubator-answer/internal/repo/answer"
	"github.com/apache/incubator-answer/internal/repo/auth"
	"github.com/apache/incubator-answer/internal/repo/badge"
	"github.com/apache/incubator-answer/internal/repo/badge_award"
	"github.com/apache/incubator-answer/internal/repo/badge_group"
	"github.com/apache/incubator-answer/internal/repo/captcha"
	"github.com/apache/incubator-answer/internal/repo/collection"
	"github.com/apache/incubator-answer/internal/repo/comment"
	"github.com/apache/incubator-answer/internal/repo/config"
	"github.com/apache/incubator-answer/internal/repo/export"
	"github.com/apache/incubator-answer/internal/repo/limit"
	"github.com/apache/incubator-answer/internal/repo/meta"
	notification2 "github.com/apache/incubator-answer/internal/repo/notification"
	"github.com/apache/incubator-answer/internal/repo/plugin_config"
	"github.com/apache/incubator-answer/internal/repo/question"
	"github.com/apache/incubator-answer/internal/repo/rank"
	"github.com/apache/incubator-answer/internal/repo/reason"
	"github.com/apache/incubator-answer/internal/repo/report"
	"github.com/apache/incubator-answer/internal/repo/review"
	"github.com/apache/incubator-answer/internal/repo/revision"
	"github.com/apache/incubator-answer/internal/repo/role"
	"github.com/apache/incubator-answer/internal/repo/search_common"
	"github.com/apache/incubator-answer/internal/repo/site_info"
	"github.com/apache/incubator-answer/internal/repo/tag"
	"github.com/apache/incubator-answer/internal/repo/tag_common"
	"github.com/apache/incubator-answer/internal/repo/unique"
	"github.com/apache/incubator-answer/internal/repo/user"
	"github.com/apache/incubator-answer/internal/repo/user_external_login"
	"github.com/apache/incubator-answer/internal/repo/user_notification_config"
	"github.com/apache/incubator-answer/internal/router"
	"github.com/apache/incubator-answer/internal/service/action"
	activity2 "github.com/apache/incubator-answer/internal/service/activity"
	activity_common2 "github.com/apache/incubator-answer/internal/service/activity_common"
	"github.com/apache/incubator-answer/internal/service/activity_queue"
	"github.com/apache/incubator-answer/internal/service/answer_common"
	auth2 "github.com/apache/incubator-answer/internal/service/auth"
	badge2 "github.com/apache/incubator-answer/internal/service/badge"
	collection2 "github.com/apache/incubator-answer/internal/service/collection"
	"github.com/apache/incubator-answer/internal/service/collection_common"
	comment2 "github.com/apache/incubator-answer/internal/service/comment"
	"github.com/apache/incubator-answer/internal/service/comment_common"
	config2 "github.com/apache/incubator-answer/internal/service/config"
	"github.com/apache/incubator-answer/internal/service/content"
	"github.com/apache/incubator-answer/internal/service/dashboard"
	"github.com/apache/incubator-answer/internal/service/event_queue"
	export2 "github.com/apache/incubator-answer/internal/service/export"
	"github.com/apache/incubator-answer/internal/service/follow"
	meta2 "github.com/apache/incubator-answer/internal/service/meta"
	"github.com/apache/incubator-answer/internal/service/meta_common"
	"github.com/apache/incubator-answer/internal/service/notice_queue"
	"github.com/apache/incubator-answer/internal/service/notification"
	"github.com/apache/incubator-answer/internal/service/notification_common"
	"github.com/apache/incubator-answer/internal/service/object_info"
	"github.com/apache/incubator-answer/internal/service/plugin_common"
	"github.com/apache/incubator-answer/internal/service/question_common"
	rank2 "github.com/apache/incubator-answer/internal/service/rank"
	reason2 "github.com/apache/incubator-answer/internal/service/reason"
	report2 "github.com/apache/incubator-answer/internal/service/report"
	"github.com/apache/incubator-answer/internal/service/report_handle"
	review2 "github.com/apache/incubator-answer/internal/service/review"
	"github.com/apache/incubator-answer/internal/service/revision_common"
	role2 "github.com/apache/incubator-answer/internal/service/role"
	"github.com/apache/incubator-answer/internal/service/search_parser"
	"github.com/apache/incubator-answer/internal/service/service_config"
	"github.com/apache/incubator-answer/internal/service/siteinfo"
	"github.com/apache/incubator-answer/internal/service/siteinfo_common"
	tag2 "github.com/apache/incubator-answer/internal/service/tag"
	tag_common2 "github.com/apache/incubator-answer/internal/service/tag_common"
	"github.com/apache/incubator-answer/internal/service/uploader"
	"github.com/apache/incubator-answer/internal/service/user_admin"
	"github.com/apache/incubator-answer/internal/service/user_common"
	user_external_login2 "github.com/apache/incubator-answer/internal/service/user_external_login"
	user_notification_config2 "github.com/apache/incubator-answer/internal/service/user_notification_config"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/log"
)

// Injectors from wire.go:

// initApplication init application.
func initApplication(debug bool, serverConf *conf.Server, dbConf *data.Database, cacheConf *data.CacheConf, i18nConf *translator.I18n, swaggerConf *router.SwaggerConfig, serviceConf *service_config.ServiceConfig, uiConf *server.UI, logConf log.Logger) (*pacman.Application, func(), error) {
	staticRouter := router.NewStaticRouter(serviceConf)
	i18nTranslator, err := translator.NewTranslator(i18nConf)
	if err != nil {
		return nil, nil, err
	}
	engine, err := data.NewDB(debug, dbConf)
	if err != nil {
		return nil, nil, err
	}
	cache, cleanup, err := data.NewCache(cacheConf)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup2, err := data.NewData(engine, cache)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	siteInfoRepo := site_info.NewSiteInfo(dataData)
	siteInfoCommonService := siteinfo_common.NewSiteInfoCommonService(siteInfoRepo)
	langController := controller.NewLangController(i18nTranslator, siteInfoCommonService)
	authRepo := auth.NewAuthRepo(dataData)
	authService := auth2.NewAuthService(authRepo)
	userRepo := user.NewUserRepo(dataData)
	uniqueIDRepo := unique.NewUniqueIDRepo(dataData)
	configRepo := config.NewConfigRepo(dataData)
	configService := config2.NewConfigService(configRepo)
	activityRepo := activity_common.NewActivityRepo(dataData, uniqueIDRepo, configService)
	userRankRepo := rank.NewUserRankRepo(dataData, configService)
	userActiveActivityRepo := activity.NewUserActiveActivityRepo(dataData, activityRepo, userRankRepo, configService)
	emailRepo := export.NewEmailRepo(dataData)
	emailService := export2.NewEmailService(configService, emailRepo, siteInfoCommonService)
	userRoleRelRepo := role.NewUserRoleRelRepo(dataData)
	roleRepo := role.NewRoleRepo(dataData)
	roleService := role2.NewRoleService(roleRepo)
	userRoleRelService := role2.NewUserRoleRelService(userRoleRelRepo, roleService)
	userCommon := usercommon.NewUserCommon(userRepo, userRoleRelService, authService, siteInfoCommonService)
	userExternalLoginRepo := user_external_login.NewUserExternalLoginRepo(dataData)
	userNotificationConfigRepo := user_notification_config.NewUserNotificationConfigRepo(dataData)
	userNotificationConfigService := user_notification_config2.NewUserNotificationConfigService(userRepo, userNotificationConfigRepo)
	userExternalLoginService := user_external_login2.NewUserExternalLoginService(userRepo, userCommon, userExternalLoginRepo, emailService, siteInfoCommonService, userActiveActivityRepo, userNotificationConfigService)
	questionRepo := question.NewQuestionRepo(dataData, uniqueIDRepo)
	answerRepo := answer.NewAnswerRepo(dataData, uniqueIDRepo, userRankRepo, activityRepo)
	voteRepo := activity_common.NewVoteRepo(dataData, activityRepo)
	followRepo := activity_common.NewFollowRepo(dataData, uniqueIDRepo, activityRepo)
	tagCommonRepo := tag_common.NewTagCommonRepo(dataData, uniqueIDRepo)
	tagRelRepo := tag.NewTagRelRepo(dataData, uniqueIDRepo)
	tagRepo := tag.NewTagRepo(dataData, uniqueIDRepo)
	revisionRepo := revision.NewRevisionRepo(dataData, uniqueIDRepo)
	revisionService := revision_common.NewRevisionService(revisionRepo, userRepo)
	activityQueueService := activity_queue.NewActivityQueueService()
	tagCommonService := tag_common2.NewTagCommonService(tagCommonRepo, tagRelRepo, tagRepo, revisionService, siteInfoCommonService, activityQueueService)
	collectionRepo := collection.NewCollectionRepo(dataData, uniqueIDRepo)
	collectionCommon := collectioncommon.NewCollectionCommon(collectionRepo)
	answerCommon := answercommon.NewAnswerCommon(answerRepo)
	metaRepo := meta.NewMetaRepo(dataData)
	metaCommonService := metacommon.NewMetaCommonService(metaRepo)
	questionCommon := questioncommon.NewQuestionCommon(questionRepo, answerRepo, voteRepo, followRepo, tagCommonService, userCommon, collectionCommon, answerCommon, metaCommonService, configService, activityQueueService, revisionRepo, dataData)
	eventQueueService := event_queue.NewEventQueueService()
	userService := content.NewUserService(userRepo, userActiveActivityRepo, activityRepo, emailService, authService, siteInfoCommonService, userRoleRelService, userCommon, userExternalLoginService, userNotificationConfigRepo, userNotificationConfigService, questionCommon, eventQueueService)
	captchaRepo := captcha.NewCaptchaRepo(dataData)
	captchaService := action.NewCaptchaService(captchaRepo)
	userController := controller.NewUserController(authService, userService, captchaService, emailService, siteInfoCommonService, userNotificationConfigService)
	commentRepo := comment.NewCommentRepo(dataData, uniqueIDRepo)
	commentCommonRepo := comment.NewCommentCommonRepo(dataData, uniqueIDRepo)
	objService := object_info.NewObjService(answerRepo, questionRepo, commentCommonRepo, tagCommonRepo, tagCommonService)
	notificationQueueService := notice_queue.NewNotificationQueueService()
	externalNotificationQueueService := notice_queue.NewNewQuestionNotificationQueueService()
	commentService := comment2.NewCommentService(commentRepo, commentCommonRepo, userCommon, objService, voteRepo, emailService, userRepo, notificationQueueService, externalNotificationQueueService, activityQueueService, eventQueueService)
	rolePowerRelRepo := role.NewRolePowerRelRepo(dataData)
	rolePowerRelService := role2.NewRolePowerRelService(rolePowerRelRepo, userRoleRelService)
	rankService := rank2.NewRankService(userCommon, userRankRepo, objService, userRoleRelService, rolePowerRelService, configService)
	limitRepo := limit.NewRateLimitRepo(dataData)
	rateLimitMiddleware := middleware.NewRateLimitMiddleware(limitRepo)
	commentController := controller.NewCommentController(commentService, rankService, captchaService, rateLimitMiddleware)
	reportRepo := report.NewReportRepo(dataData, uniqueIDRepo)
	tagService := tag2.NewTagService(tagRepo, tagCommonService, revisionService, followRepo, siteInfoCommonService, activityQueueService)
	answerActivityRepo := activity.NewAnswerActivityRepo(dataData, activityRepo, userRankRepo, notificationQueueService)
	answerActivityService := activity2.NewAnswerActivityService(answerActivityRepo, configService)
	externalNotificationService := notification.NewExternalNotificationService(dataData, userNotificationConfigRepo, followRepo, emailService, userRepo, externalNotificationQueueService, userExternalLoginRepo, siteInfoCommonService)
	reviewRepo := review.NewReviewRepo(dataData)
	reviewService := review2.NewReviewService(reviewRepo, objService, userCommon, userRepo, questionRepo, answerRepo, userRoleRelService, externalNotificationQueueService, tagCommonService, questionCommon, notificationQueueService, siteInfoCommonService)
	questionService := content.NewQuestionService(activityRepo, questionRepo, answerRepo, tagCommonService, tagService, questionCommon, userCommon, userRepo, userRoleRelService, revisionService, metaCommonService, collectionCommon, answerActivityService, emailService, notificationQueueService, externalNotificationQueueService, activityQueueService, siteInfoCommonService, externalNotificationService, reviewService, configService, eventQueueService)
	answerService := content.NewAnswerService(answerRepo, questionRepo, questionCommon, userCommon, collectionCommon, userRepo, revisionService, answerActivityService, answerCommon, voteRepo, emailService, userRoleRelService, notificationQueueService, externalNotificationQueueService, activityQueueService, reviewService, eventQueueService)
	reportHandle := report_handle.NewReportHandle(questionService, answerService, commentService)
	reportService := report2.NewReportService(reportRepo, objService, userCommon, answerRepo, questionRepo, commentCommonRepo, reportHandle, configService, eventQueueService)
	reportController := controller.NewReportController(reportService, rankService, captchaService)
	contentVoteRepo := activity.NewVoteRepo(dataData, activityRepo, userRankRepo, notificationQueueService)
	voteService := content.NewVoteService(contentVoteRepo, configService, questionRepo, answerRepo, commentCommonRepo, objService, eventQueueService)
	voteController := controller.NewVoteController(voteService, rankService, captchaService)
	tagController := controller.NewTagController(tagService, tagCommonService, rankService)
	followFollowRepo := activity.NewFollowRepo(dataData, uniqueIDRepo, activityRepo)
	followService := follow.NewFollowService(followFollowRepo, followRepo, tagCommonRepo)
	followController := controller.NewFollowController(followService)
	collectionGroupRepo := collection.NewCollectionGroupRepo(dataData)
	collectionService := collection2.NewCollectionService(collectionRepo, collectionGroupRepo, questionCommon)
	collectionController := controller.NewCollectionController(collectionService)
	questionController := controller.NewQuestionController(questionService, answerService, rankService, siteInfoCommonService, captchaService, rateLimitMiddleware)
	answerController := controller.NewAnswerController(answerService, rankService, captchaService, siteInfoCommonService, rateLimitMiddleware)
	searchParser := search_parser.NewSearchParser(tagCommonService, userCommon)
	searchRepo := search_common.NewSearchRepo(dataData, uniqueIDRepo, userCommon, tagCommonService)
	searchService := content.NewSearchService(searchParser, searchRepo)
	searchController := controller.NewSearchController(searchService, captchaService)
	reviewActivityRepo := activity.NewReviewActivityRepo(dataData, activityRepo, userRankRepo, configService)
	contentRevisionService := content.NewRevisionService(revisionRepo, userCommon, questionCommon, answerService, objService, questionRepo, answerRepo, tagRepo, tagCommonService, notificationQueueService, activityQueueService, reportRepo, reviewService, reviewActivityRepo)
	revisionController := controller.NewRevisionController(contentRevisionService, rankService)
	rankController := controller.NewRankController(rankService)
	userAdminRepo := user.NewUserAdminRepo(dataData, authRepo)
	userAdminService := user_admin.NewUserAdminService(userAdminRepo, userRoleRelService, authService, userCommon, userActiveActivityRepo, siteInfoCommonService, emailService, questionRepo, answerRepo, commentCommonRepo)
	userAdminController := controller_admin.NewUserAdminController(userAdminService)
	reasonRepo := reason.NewReasonRepo(configService)
	reasonService := reason2.NewReasonService(reasonRepo)
	reasonController := controller.NewReasonController(reasonService)
	themeController := controller_admin.NewThemeController()
	siteInfoService := siteinfo.NewSiteInfoService(siteInfoRepo, siteInfoCommonService, emailService, tagCommonService, configService, questionCommon)
	siteInfoController := controller_admin.NewSiteInfoController(siteInfoService)
	controllerSiteInfoController := controller.NewSiteInfoController(siteInfoCommonService)
	notificationRepo := notification2.NewNotificationRepo(dataData)
	notificationCommon := notificationcommon.NewNotificationCommon(dataData, notificationRepo, userCommon, activityRepo, followRepo, objService, notificationQueueService, userExternalLoginRepo, siteInfoCommonService)
	badgeRepo := badge.NewBadgeRepo(dataData, uniqueIDRepo)
	notificationService := notification.NewNotificationService(dataData, notificationRepo, notificationCommon, revisionService, userRepo, reportRepo, reviewService, badgeRepo)
	notificationController := controller.NewNotificationController(notificationService, rankService)
	dashboardService := dashboard.NewDashboardService(questionRepo, answerRepo, commentCommonRepo, voteRepo, userRepo, reportRepo, configService, siteInfoCommonService, serviceConf, reviewService, revisionRepo, dataData)
	dashboardController := controller.NewDashboardController(dashboardService)
	uploaderService := uploader.NewUploaderService(serviceConf, siteInfoCommonService)
	uploadController := controller.NewUploadController(uploaderService)
	activityActivityRepo := activity.NewActivityRepo(dataData, configService)
	activityCommon := activity_common2.NewActivityCommon(activityRepo, activityQueueService)
	commentCommonService := comment_common.NewCommentCommonService(commentCommonRepo)
	activityService := activity2.NewActivityService(activityActivityRepo, userCommon, activityCommon, tagCommonService, objService, commentCommonService, revisionService, metaCommonService, configService)
	activityController := controller.NewActivityController(activityService)
	roleController := controller_admin.NewRoleController(roleService)
	pluginConfigRepo := plugin_config.NewPluginConfigRepo(dataData)
	pluginUserConfigRepo := plugin_config.NewPluginUserConfigRepo(dataData)
	pluginCommonService := plugin_common.NewPluginCommonService(pluginConfigRepo, pluginUserConfigRepo, configService, dataData)
	pluginController := controller_admin.NewPluginController(pluginCommonService)
	permissionController := controller.NewPermissionController(rankService)
	userPluginController := controller.NewUserPluginController(pluginCommonService)
	reviewController := controller.NewReviewController(reviewService, rankService, captchaService)
	metaService := meta2.NewMetaService(metaCommonService, userCommon, answerRepo, questionRepo, eventQueueService)
	metaController := controller.NewMetaController(metaService)
	badgeGroupRepo := badge_group.NewBadgeGroupRepo(dataData, uniqueIDRepo)
	badgeAwardRepo := badge_award.NewBadgeAwardRepo(dataData, uniqueIDRepo)
	eventRuleRepo := badge.NewEventRuleRepo(dataData)
	badgeAwardService := badge2.NewBadgeAwardService(badgeAwardRepo, badgeRepo, userCommon, objService, notificationQueueService)
	badgeEventService := badge2.NewBadgeEventService(dataData, eventQueueService, badgeRepo, eventRuleRepo, badgeAwardService)
	badgeService := badge2.NewBadgeService(badgeRepo, badgeGroupRepo, badgeAwardRepo, badgeEventService, siteInfoCommonService)
	badgeController := controller.NewBadgeController(badgeService, badgeAwardService)
	controller_adminBadgeController := controller_admin.NewBadgeController(badgeService)
	answerAPIRouter := router.NewAnswerAPIRouter(langController, userController, commentController, reportController, voteController, tagController, followController, collectionController, questionController, answerController, searchController, revisionController, rankController, userAdminController, reasonController, themeController, siteInfoController, controllerSiteInfoController, notificationController, dashboardController, uploadController, activityController, roleController, pluginController, permissionController, userPluginController, reviewController, metaController, badgeController, controller_adminBadgeController)
	swaggerRouter := router.NewSwaggerRouter(swaggerConf)
	uiRouter := router.NewUIRouter(controllerSiteInfoController, siteInfoCommonService)
	authUserMiddleware := middleware.NewAuthUserMiddleware(authService, siteInfoCommonService)
	avatarMiddleware := middleware.NewAvatarMiddleware(serviceConf, uploaderService)
	shortIDMiddleware := middleware.NewShortIDMiddleware(siteInfoCommonService)
	templateRenderController := templaterender.NewTemplateRenderController(questionService, userService, tagService, answerService, commentService, siteInfoCommonService, questionRepo)
	templateController := controller.NewTemplateController(templateRenderController, siteInfoCommonService, eventQueueService, userService)
	templateRouter := router.NewTemplateRouter(templateController, templateRenderController, siteInfoController, authUserMiddleware)
	connectorController := controller.NewConnectorController(siteInfoCommonService, emailService, userExternalLoginService)
	userCenterLoginService := user_external_login2.NewUserCenterLoginService(userRepo, userCommon, userExternalLoginRepo, userActiveActivityRepo, siteInfoCommonService)
	userCenterController := controller.NewUserCenterController(userCenterLoginService, siteInfoCommonService)
	captchaController := controller.NewCaptchaController()
	embedController := controller.NewEmbedController()
	renderController := controller.NewRenderController()
	importerController := controller.NewImporterController(questionService, rankService, rateLimitMiddleware, userCommon)
	pluginAPIRouter := router.NewPluginAPIRouter(connectorController, userCenterController, captchaController, embedController, renderController, importerController)
	ginEngine := server.NewHTTPServer(debug, staticRouter, answerAPIRouter, swaggerRouter, uiRouter, authUserMiddleware, avatarMiddleware, shortIDMiddleware, templateRouter, pluginAPIRouter, uiConf)
	scheduledTaskManager := cron.NewScheduledTaskManager(siteInfoCommonService, questionService)
	application := newApplication(serverConf, ginEngine, scheduledTaskManager)
	return application, func() {
		cleanup2()
		cleanup()
	}, nil
}
