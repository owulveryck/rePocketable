package main

import "time"

type Toc struct {
	Environment struct {
		Origin        string `json:"origin"`
		Env           string `json:"env"`
		Nodeenv       string `json:"nodeEnv"`
		Frontendpaths struct {
			Books struct {
				Anonymoushostname string `json:"anonymousHostname"`
				Hostname          string `json:"hostname"`
				Path              string `json:"path"`
				Routes            []struct {
					Contenttype string   `json:"contentType"`
					Type        string   `json:"type"`
					Path        string   `json:"path"`
					Exact       bool     `json:"exact"`
					Parameters  []string `json:"parameters"`
				} `json:"routes"`
			} `json:"books"`
			Learningpaths struct {
				Anonymoushostname string `json:"anonymousHostname"`
				Hostname          string `json:"hostname"`
				Path              string `json:"path"`
				Routes            []struct {
					Contenttype  string   `json:"contentType"`
					Type         string   `json:"type"`
					Path         string   `json:"path"`
					Exact        bool     `json:"exact"`
					Parameters   []string `json:"parameters"`
					Searchparams struct {
						SnippetID string `json:"snippet_id"`
					} `json:"searchParams,omitempty"`
				} `json:"routes"`
			} `json:"learningPaths"`
			Videos struct {
				Anonymoushostname string `json:"anonymousHostname"`
				Anonymouspath     string `json:"anonymousPath"`
				Hostname          string `json:"hostname"`
				Path              string `json:"path"`
				Routes            []struct {
					Contenttype   string   `json:"contentType"`
					Type          string   `json:"type"`
					Path          string   `json:"path"`
					Exact         bool     `json:"exact"`
					Parameters    []string `json:"parameters"`
					Anonymouspath string   `json:"anonymousPath,omitempty"`
				} `json:"routes"`
			} `json:"videos"`
			Scenarios struct {
				Hostname string `json:"hostname"`
				Path     string `json:"path"`
				Routes   []struct {
					Contenttype string   `json:"contentType"`
					Type        string   `json:"type"`
					Path        string   `json:"path"`
					Exact       bool     `json:"exact"`
					Parameters  []string `json:"parameters"`
				} `json:"routes"`
			} `json:"scenarios"`
		} `json:"frontendPaths"`
		Frontendname          string   `json:"frontendName"`
		Tld                   string   `json:"tld"`
		Xstatedevtoolsenabled bool     `json:"xstateDevToolsEnabled"`
		Learningurl           string   `json:"learningUrl"`
		Apipath               string   `json:"apiPath"`
		Ingressedpaths        []string `json:"ingressedPaths"`
		Basepath              string   `json:"basePath"`
		Publicpath            string   `json:"publicPath"`
		Routerbasepath        string   `json:"routerBasePath"`
	} `json:"environment"`
	Ldfeatureflags struct {
		Loadingstate                               string      `json:"_loadingState"`
		Error                                      interface{} `json:"_error"`
		Abadflagname                               bool        `json:"aBadFlagname"`
		Accountpersonalinfo                        bool        `json:"accountPersonalInfo"`
		Alleventspage                              bool        `json:"allEventsPage"`
		Allowreaderpagingios                       bool        `json:"allowReaderPagingIos"`
		Androidenableamplitude                     bool        `json:"androidEnableAmplitude"`
		Androidenablefullstory                     bool        `json:"androidEnableFullstory"`
		Androidenablemarketingcloudsdk             bool        `json:"androidEnableMarketingCloudSdk"`
		Answersexcludevideo                        bool        `json:"answersExcludeVideo"`
		Askforreviewsios                           bool        `json:"askForReviewsIos"`
		Assignmentsbffgraphql                      bool        `json:"assignmentsBffGraphql"`
		Attendsponsorships                         bool        `json:"attendSponsorships"`
		Audiobooksinsightdashboardenabled          bool        `json:"audiobooksInsightDashboardEnabled"`
		Betasegmentnote                            bool        `json:"betaSegmentNote"`
		Boostforauthorandpublisherqueries          bool        `json:"boostForAuthorAndPublisherQueries"`
		Cachesearchresultsinsessionstorage         bool        `json:"cacheSearchResultsInSessionStorage"`
		Certificationshidepracticeexams            bool        `json:"certificationsHidePracticeExams"`
		Certificationspartneroutage                bool        `json:"certificationsPartnerOutage"`
		Certificationsshowpelink                   bool        `json:"certificationsShowPeLink"`
		Chassishellotoggle                         bool        `json:"chassisHelloToggle"`
		Classifyliveeventseriesfromcontentmessages bool        `json:"classifyLiveEventSeriesFromContentMessages"`
		Conferencesdroid                           bool        `json:"conferencesDroid"`
		Contentplaygroundprototype                 bool        `json:"contentPlaygroundPrototype"`
		Contentsearchexperiment                    bool        `json:"contentSearchExperiment"`
		Copyhighlightlink                          bool        `json:"copyHighlightLink"`
		Disableingestionportal                     struct {
			DisableFtpProcessing bool `json:"disable_ftp_processing"`
			DisablePortal        bool `json:"disable_portal"`
		} `json:"disableIngestionPortal"`
		Disablelocaltrainingnotificationsios            bool     `json:"disableLocalTrainingNotificationsIos"`
		Disableviewnotebook                             bool     `json:"disableViewNotebook"`
		Displayaudiobooks                               bool     `json:"displayAudiobooks"`
		Displayjalapenopoppers                          bool     `json:"displayJalapenoPoppers"`
		Displaynewemailpreferences                      bool     `json:"displayNewEmailPreferences"`
		Displayonetrustembed                            bool     `json:"displayOneTrustEmbed"`
		Displayverticalfilters                          bool     `json:"displayVerticalFilters"`
		Dsexperiencetest                                string   `json:"dsExperienceTest"`
		Eisdisablesampleuploads                         bool     `json:"eisDisableSampleUploads"`
		Enableamplitudeios                              bool     `json:"enableAmplitudeIos"`
		Enableanswersclient                             bool     `json:"enableAnswersClient"`
		Enablebookmarks                                 bool     `json:"enableBookmarks"`
		Enablecloudscenarios                            bool     `json:"enableCloudScenarios"`
		Enablecontentplaygroundfrontend                 bool     `json:"enableContentPlaygroundFrontend"`
		Enabledocumentclient                            bool     `json:"enableDocumentClient"`
		Enableepubdiffprocessing                        bool     `json:"enableEpubDiffProcessing"`
		Enablefeatherheronauth                          bool     `json:"enableFeatherHeronAuth"`
		Enablefeedbacklink                              bool     `json:"enableFeedbackLink"`
		Enablefsonsearch                                bool     `json:"enableFsOnSearch"`
		Enablefullaudiobooksexperience                  bool     `json:"enableFullAudiobooksExperience"`
		Enablefullstoryios                              bool     `json:"enableFullstoryIos"`
		Enablehomepage                                  bool     `json:"enableHomepage"`
		Enableinvitationsfe                             bool     `json:"enableInvitationsFe"`
		Enablejonestest                                 bool     `json:"enableJonesTest"`
		Enablelargemessagesvmsingestion                 bool     `json:"enableLargeMessagesVmsIngestion"`
		Enablemagpieimagesoncerts                       bool     `json:"enableMagpieImagesOnCerts"`
		Enablemedsgrootetsync                           bool     `json:"enableMedsGrootEtSync"`
		Enablenewregisterpage                           bool     `json:"enableNewRegisterPage"`
		Enablenewsignuppage                             bool     `json:"enableNewSignupPage"`
		Enablenewsubscribepage                          bool     `json:"enableNewSubscribePage"`
		Enableoptimizemultiselect                       bool     `json:"enableOptimizeMultiselect"`
		Enableormchromeextension                        bool     `json:"enableOrmChromeExtension"`
		Enablepausesubscriptionpage                     bool     `json:"enablePauseSubscriptionPage"`
		Enablepaypalchoicebuttoninnewpaymentsclient     bool     `json:"enablePayPalChoiceButtonInNewPaymentsClient"`
		Enablepaymentspricetiers                        bool     `json:"enablePaymentsPriceTiers"`
		Enablepdsv1Alphacontentingestionmessages        bool     `json:"enablePdsV1AlphaContentIngestionMessages"`
		Enablepdsv1Alphaproductingestionmessages        bool     `json:"enablePdsV1AlphaProductIngestionMessages"`
		Enablerateplancallouts                          bool     `json:"enableRatePlanCallouts"`
		Enablerokufiretv                                bool     `json:"enableRokuFiretv"`
		Enablesemiannualplan                            bool     `json:"enableSemiAnnualPlan"`
		Enablesignup                                    bool     `json:"enableSignup"`
		Enablespeedreaderinsearch                       bool     `json:"enableSpeedReaderInSearch"`
		Enableteamsetuppage                             bool     `json:"enableTeamSetupPage"`
		Enablethreemonthplan                            bool     `json:"enableThreeMonthPlan"`
		Enabletimebasedusageandroid                     bool     `json:"enableTimeBasedUsageAndroid"`
		Enabletimebasedusageios                         bool     `json:"enableTimeBasedUsageIos"`
		Enabletopicgraphv1Alphacontentingestionmessages bool     `json:"enableTopicGraphV1AlphaContentIngestionMessages"`
		Enabletvapps                                    bool     `json:"enableTvApps"`
		Enableusermanagement                            bool     `json:"enableUserManagement"`
		Enableusermanagementclient                      bool     `json:"enableUserManagementClient"`
		Enablevideocourses                              bool     `json:"enableVideoCourses"`
		Enablewebsite                                   bool     `json:"enableWebsite"`
		Enableworkbooks                                 bool     `json:"enableWorkbooks"`
		Enableyouraccountlink                           bool     `json:"enableYourAccountLink"`
		Enabledthirdpartyauthproviders                  []string `json:"enabledThirdPartyAuthProviders"`
		Excludeclosedliveevents                         bool     `json:"excludeClosedLiveEvents"`
		Falconcisurls                                   bool     `json:"falconCisUrls"`
		Falconcisurlsforlp                              bool     `json:"falconCisUrlsForLp"`
		Falconlprefidurl                                bool     `json:"falconLpRefidUrl"`
		Fechassishelloenableconfirm                     bool     `json:"feChassisHelloEnableConfirm"`
		Focussearchbarshortcut                          bool     `json:"focusSearchBarShortcut"`
		Heronhandlespaymenterrors                       bool     `json:"heronHandlesPaymentErrors"`
		Heronreaderignoreusageiscomplete                bool     `json:"heronReaderIgnoreUsageIsComplete"`
		Homeenableanswersanimation                      bool     `json:"homeEnableAnswersAnimation"`
		Interactivecontent                              bool     `json:"interactiveContent"`
		Interactivepracticeexamdropdownfilterenabled    bool     `json:"interactivePracticeExamDropdownFilterEnabled"`
		Jupyternotebookinsearch                         bool     `json:"jupyterNotebookInSearch"`
		Jupyterv1Web                                    bool     `json:"jupyterV1Web"`
		Labsmvptemplate                                 bool     `json:"labsMvpTemplate"`
		Labsormstaff                                    bool     `json:"labsOrmStaff"`
		Learningpathsdroid                              bool     `json:"learningPathsDroid"`
		Liveeventappearinrecentlyadded                  bool     `json:"liveEventAppearInRecentlyAdded"`
		Liveeventdetailspage                            bool     `json:"liveEventDetailsPage"`
		Liveeventfullproductmessageingestion            bool     `json:"liveEventFullProductMessageIngestion"`
		Liveeventingestfromproductmessage               bool     `json:"liveEventIngestFromProductMessage"`
		Liveeventrecurringregistration                  bool     `json:"liveEventRecurringRegistration"`
		Liveeventsendcontentmessages                    bool     `json:"liveEventSendContentMessages"`
		Liveeventseriespage                             bool     `json:"liveEventSeriesPage"`
		Liveeventsponsorshipemails                      bool     `json:"liveEventSponsorshipEmails"`
		Liveeventuserregconsumecontentmessages          bool     `json:"liveEventUserRegConsumeContentMessages"`
		Liveeventscardenhancements                      bool     `json:"liveEventsCardEnhancements"`
		Liveeventsenablecaching                         bool     `json:"liveEventsEnableCaching"`
		Liveeventstimezonefilter                        bool     `json:"liveEventsTimezoneFilter"`
		Loglowrecaptchascores                           float64  `json:"logLowRecaptchaScores"`
		Manytomanyssodevtools                           bool     `json:"manyToManySsoDevTools"`
		Marcshowbanner                                  bool     `json:"marcShowBanner"`
		Moresearchresulttypes                           bool     `json:"moreSearchResultTypes"`
		Multilateralfederatedauthintegrationanybird     bool     `json:"multilateralFederatedAuthIntegrationAnybird"`
		Multilateralfederatedauthintegrationfe          bool     `json:"multilateralFederatedAuthIntegrationFe"`
		Oreillylabs                                     bool     `json:"oreillyLabs"`
		Paymentsclientshowcompareplanslink              bool     `json:"paymentsClientShowComparePlansLink"`
		Paymentsclientshowyourdeviceslink               bool     `json:"paymentsClientShowYourDevicesLink"`
		Pearsoninteractiveinside                        bool     `json:"pearsonInteractiveInside"`
		Playlistaccountsharing                          bool     `json:"playlistAccountSharing"`
		Playlistsaddsections                            bool     `json:"playlistsAddSections"`
		Practiceexamsinsearch                           bool     `json:"practiceExamsInSearch"`
		Publicplaylistsandroid                          bool     `json:"publicPlaylistsAndroid"`
		Rainbowheadlines                                bool     `json:"rainbowHeadlines"`
		Redirectcreatetrialtostarttrial                 bool     `json:"redirectCreateTrialToStartTrial"`
		Redirectlegacyoauthcompletions                  bool     `json:"redirectLegacyOauthCompletions"`
		Redirectpregistertostarttrial                   bool     `json:"redirectPRegisterToStartTrial"`
		Redirectresultswithweburl                       bool     `json:"redirectResultsWithWebUrl"`
		Sandboxscenariosinsearch                        bool     `json:"sandboxScenariosInSearch"`
		Searchbarvariation                              string   `json:"searchBarVariation"`
		Searchjudgement                                 bool     `json:"searchJudgement"`
		Searchrelevancetuning                           struct {
			Pqf struct {
				ContentUnstemmed      int `json:"content_unstemmed"`
				TitleUnstemmed        int `json:"title_unstemmed"`
				TopicKeywords         int `json:"topic_keywords"`
				TopicNames            int `json:"topic_names"`
				ChapterTitleUnstemmed int `json:"chapter_title_unstemmed"`
			} `json:"pqf"`
			Qf struct {
				Content               int `json:"content"`
				Title                 int `json:"title"`
				TopicKeywords         int `json:"topic_keywords"`
				TopicNames            int `json:"topic_names"`
				ChapterTitle          int `json:"chapter_title"`
				ChapterTitleUnstemmed int `json:"chapter_title_unstemmed"`
			} `json:"qf"`
			Tie       float64 `json:"tie"`
			Boost     string  `json:"boost"`
			Collapse  bool    `json:"collapse"`
			Highlight bool    `json:"highlight"`
		} `json:"searchRelevanceTuning"`
		Seatcountmismatchthresholds struct {
			SeatCapacityPercentage float64 `json:"seat_capacity_percentage"`
			DifferencePercentage   float64 `json:"difference_percentage"`
		} `json:"seatCountMismatchThresholds"`
		Selfregistrationrollout          bool `json:"selfRegistrationRollout"`
		Sendcanonproductmessage          bool `json:"sendCanonProductMessage"`
		Senddevicetokenstomarketingcloud bool `json:"sendDeviceTokensToMarketingCloud"`
		Sendtransactionalemails          bool `json:"sendTransactionalEmails"`
		Sfmcaccountcontactevent          bool `json:"sfmcAccountContactEvent"`
		Shouldcausecrashios              bool `json:"shouldCauseCrashIos"`
		Shoulduseunifiedhistoryios       bool `json:"shouldUseUnifiedHistoryIos"`
		Showinternaltoolsios             bool `json:"showInternalToolsIos"`
		Showlinkappleicon                bool `json:"showLinkAppleIcon"`
		Showpaymentsnotificationalert    bool `json:"showPaymentsNotificationAlert"`
		Showremainingduration            bool `json:"showRemainingDuration"`
		Showreminderbarperld             bool `json:"showReminderBarPerLd"`
		Showscenariosinsearch            bool `json:"showScenariosInSearch"`
		Showsearchresulttopic            bool `json:"showSearchResultTopic"`
		Showuserdetails                  bool `json:"showUserDetails"`
		Solrfieldalias                   bool `json:"solrFieldAlias"`
		Syncgrootpasswords               bool `json:"syncGrootPasswords"`
		Teachbeta                        bool `json:"teachBeta"`
		Topicsv2Viewcontentcounts        bool `json:"topicsV2ViewContentCounts"`
		Universalcontentviewerconfig     struct {
			LearningPaths bool `json:"/learning-paths/"`
			LibraryView   bool `json:"/library/view/"`
			Videos        bool `json:"/videos/"`
		} `json:"universalContentViewerConfig"`
		Usejsonfacetsinserp                          bool `json:"useJsonFacetsInSerp"`
		Usemultiselectfacetsinserp                   bool `json:"useMultiselectFacetsInSerp"`
		Usenewinvitationurls                         bool `json:"useNewInvitationUrls"`
		Validateinteractivityproxyurl                bool `json:"validateInteractivityProxyUrl"`
		Videointestionserviceusetranscriptionservice bool `json:"videoIntestionServiceUseTranscriptionService"`
		Visautotranscription                         struct {
			Prefixes []string `json:"prefixes"`
		} `json:"visAutoTranscription"`
		Whatisjupyternotebooklink bool `json:"whatIsJupyterNotebookLink"`
	} `json:"ldFeatureFlags"`
	Jwt struct {
		Loadingstate string      `json:"_loadingState"`
		Error        interface{} `json:"_error"`
		Accts        []string    `json:"accts"`
		Eids         struct {
			Exacttarget string `json:"exacttarget"`
			Heron       string `json:"heron"`
		} `json:"eids"`
		Env        string `json:"env"`
		Exp        int    `json:"exp"`
		Individual bool   `json:"individual"`
		Perms      struct {
			Acadm  string `json:"acadm"`
			Apidc  string `json:"apidc"`
			Asign  string `json:"asign"`
			Cldsc  string `json:"cldsc"`
			Cnfrc  string `json:"cnfrc"`
			Cprex  string `json:"cprex"`
			Csstd  string `json:"csstd"`
			Epubs  string `json:"epubs"`
			Lrpth  string `json:"lrpth"`
			Lvtrg  string `json:"lvtrg"`
			Ntbks  string `json:"ntbks"`
			Oriol  string `json:"oriol"`
			Plyls  string `json:"plyls"`
			Scnrio string `json:"scnrio"`
			Usage  string `json:"usage"`
			Usrpf  string `json:"usrpf"`
			Video  string `json:"video"`
		} `json:"perms"`
		Sub string `json:"sub"`
	} `json:"jwt"`
	Navigationandannouncements struct {
		Announcements []interface{} `json:"announcements"`
		Links         struct {
			Fineprint []struct {
				Name     string        `json:"name"`
				Link     string        `json:"link"`
				Icon     string        `json:"icon"`
				Groups   []string      `json:"groups"`
				Children []interface{} `json:"children"`
			} `json:"fineprint"`
			Navigation []struct {
				Name     string        `json:"name"`
				Link     string        `json:"link"`
				Icon     string        `json:"icon"`
				Groups   []string      `json:"groups"`
				Children []interface{} `json:"children"`
			} `json:"navigation"`
			Yourprofile []struct {
				Name     string   `json:"name"`
				Link     string   `json:"link"`
				Icon     string   `json:"icon"`
				Groups   []string `json:"groups"`
				Children []struct {
					Name string `json:"name"`
					Link string `json:"link"`
					Icon string `json:"icon"`
				} `json:"children"`
			} `json:"yourprofile"`
			Footer []struct {
				Name     string        `json:"name"`
				Link     string        `json:"link"`
				Icon     string        `json:"icon"`
				Groups   []string      `json:"groups"`
				Children []interface{} `json:"children"`
			} `json:"footer"`
			Secondary []struct {
				Name     string        `json:"name"`
				Link     string        `json:"link"`
				Icon     string        `json:"icon"`
				Groups   []string      `json:"groups"`
				Children []interface{} `json:"children"`
			} `json:"secondary"`
		} `json:"links"`
	} `json:"navigationAndAnnouncements"`
	User struct {
		Error                     interface{}   `json:"error"`
		AcademicInstitution       bool          `json:"academic_institution"`
		DateJoined                time.Time     `json:"date_joined"`
		DisablePublicSharing      bool          `json:"disable_public_sharing"`
		DotgovCompAccount         bool          `json:"dotgov_comp_account"`
		Email                     string        `json:"email"`
		FirstName                 string        `json:"first_name"`
		HighlightPrivacy          string        `json:"highlight_privacy"`
		HighlightPrivacyUpdated   time.Time     `json:"highlight_privacy_updated"`
		Individual                bool          `json:"individual"`
		IsUnified                 bool          `json:"is_unified"`
		IsTrial                   bool          `json:"is_trial"`
		LastName                  string        `json:"last_name"`
		MaxSynchronousUsageEvents int           `json:"max_synchronous_usage_events"`
		OauthConnections          []interface{} `json:"oauth_connections"`
		OrganizationName          string        `json:"organization_name"`
		Permissions               struct {
			ViewAccountReporting   bool `json:"view_account_reporting"`
			ViewTeamAdministration bool `json:"view_team_administration"`
			ViewCollections        bool `json:"view_collections"`
			ViewFullEpub           bool `json:"view_full_epub"`
		} `json:"permissions"`
		PreferredLanguages     []string    `json:"preferred_languages"`
		PrefetchPaginateBy     int         `json:"prefetch_paginate_by"`
		PrimaryAccount         string      `json:"primary_account"`
		PrimaryOauthConnection interface{} `json:"primary_oauth_connection"`
		ProfileCreatedDate     string      `json:"profile_created_date"`
		SiteStyles             []string    `json:"site_styles"`
		Subscription           struct {
			Active           bool   `json:"active"`
			CancellationDate string `json:"cancellation_date"`
			Complimentary    bool   `json:"complimentary"`
		} `json:"subscription"`
		Trial struct {
			Trial               bool        `json:"trial"`
			TrialExpirationDate interface{} `json:"trial_expiration_date"`
		} `json:"trial"`
		ExpiredTrial   bool   `json:"expired_trial"`
		Username       string `json:"username"`
		UserIdentifier string `json:"user_identifier"`
		UserType       string `json:"user_type"`
		Sessionid      string `json:"sessionid"`
	} `json:"user"`
	Userinfo struct {
		Meta struct {
			Error                     interface{}   `json:"error"`
			AcademicInstitution       bool          `json:"academic_institution"`
			DateJoined                time.Time     `json:"date_joined"`
			DisablePublicSharing      bool          `json:"disable_public_sharing"`
			DotgovCompAccount         bool          `json:"dotgov_comp_account"`
			Email                     string        `json:"email"`
			FirstName                 string        `json:"first_name"`
			HighlightPrivacy          string        `json:"highlight_privacy"`
			HighlightPrivacyUpdated   time.Time     `json:"highlight_privacy_updated"`
			Individual                bool          `json:"individual"`
			IsUnified                 bool          `json:"is_unified"`
			IsTrial                   bool          `json:"is_trial"`
			LastName                  string        `json:"last_name"`
			MaxSynchronousUsageEvents int           `json:"max_synchronous_usage_events"`
			OauthConnections          []interface{} `json:"oauth_connections"`
			OrganizationName          string        `json:"organization_name"`
			Permissions               struct {
				ViewAccountReporting   bool `json:"view_account_reporting"`
				ViewTeamAdministration bool `json:"view_team_administration"`
				ViewCollections        bool `json:"view_collections"`
				ViewFullEpub           bool `json:"view_full_epub"`
			} `json:"permissions"`
			PreferredLanguages     []string    `json:"preferred_languages"`
			PrefetchPaginateBy     int         `json:"prefetch_paginate_by"`
			PrimaryAccount         string      `json:"primary_account"`
			PrimaryOauthConnection interface{} `json:"primary_oauth_connection"`
			ProfileCreatedDate     string      `json:"profile_created_date"`
			SiteStyles             []string    `json:"site_styles"`
			Subscription           struct {
				Active           bool   `json:"active"`
				CancellationDate string `json:"cancellation_date"`
				Complimentary    bool   `json:"complimentary"`
			} `json:"subscription"`
			Trial struct {
				Trial               bool        `json:"trial"`
				TrialExpirationDate interface{} `json:"trial_expiration_date"`
			} `json:"trial"`
			ExpiredTrial   bool   `json:"expired_trial"`
			Username       string `json:"username"`
			UserIdentifier string `json:"user_identifier"`
			UserType       string `json:"user_type"`
			Sessionid      string `json:"sessionid"`
		} `json:"meta"`
		Profile struct {
			Error interface{} `json:"error"`
		} `json:"profile"`
		Usagestatus struct {
			Error       interface{} `json:"error"`
			Usagestatus bool        `json:"usageStatus"`
			Success     bool        `json:"success"`
		} `json:"usageStatus"`
	} `json:"userInfo"`
	Userprofile struct {
		Loadingstate string `json:"_loadingState"`
	} `json:"userProfile"`
	Playlistscorestate struct {
		Mostrecentplaylistids []interface{} `json:"mostRecentPlaylistIds"`
		Playlists             struct {
			Error          string        `json:"error"`
			Etag           string        `json:"eTag"`
			Fetching       bool          `json:"fetching"`
			Loaded         bool          `json:"loaded"`
			Playlists      []interface{} `json:"playlists"`
			Sharingenabled bool          `json:"sharingEnabled"`
		} `json:"playlists"`
		Playlistactions struct {
			Updatingsharing bool `json:"updatingSharing"`
			Sharingerrors   struct {
			} `json:"sharingErrors"`
		} `json:"playlistActions"`
		Playlistssr struct {
			Error    string `json:"error"`
			Fetching bool   `json:"fetching"`
			Loaded   bool   `json:"loaded"`
			Playlist struct {
			} `json:"playlist"`
		} `json:"playlistSSR"`
		Sharedplaylist struct {
			Error    string `json:"error"`
			Fetching bool   `json:"fetching"`
			Loaded   bool   `json:"loaded"`
			Playlist struct {
			} `json:"playlist"`
		} `json:"sharedPlaylist"`
	} `json:"playlistsCoreState"`
	Chapterreader struct {
		Content struct {
			Chapters struct {
			} `json:"chapters"`
			Errors struct {
			} `json:"errors"`
		} `json:"content"`
		Progress struct {
			Chapters struct {
			} `json:"chapters"`
			Errors struct {
			} `json:"errors"`
		} `json:"progress"`
	} `json:"chapterReader"`
	Reviewsstate struct {
		Batchreports struct {
		} `json:"batchReports"`
		Titlereports struct {
		} `json:"titleReports"`
		Userreport struct {
		} `json:"userReport"`
		Modal struct {
			Ismounted bool `json:"isMounted"`
			Editmode  bool `json:"editMode"`
		} `json:"modal"`
		Notification struct {
			Active  bool   `json:"active"`
			Message string `json:"message"`
		} `json:"notification"`
		Loadingstates struct {
			Batch string `json:"batch"`
			Title string `json:"title"`
			User  string `json:"user"`
		} `json:"loadingStates"`
		Error string `json:"error"`
	} `json:"reviewsState"`
	Videoplayer struct {
		Content struct {
		} `json:"content"`
		Kaltura struct {
			Kalturasession struct {
				Loadingstate string      `json:"_loadingState"`
				Expiry       interface{} `json:"expiry"`
				Privileges   interface{} `json:"privileges"`
				Session      interface{} `json:"session"`
			} `json:"kalturaSession"`
			Playersettings struct {
				Playbackrate int `json:"playbackRate"`
			} `json:"playerSettings"`
		} `json:"kaltura"`
		Platformsettings struct {
			Loadingstate  string `json:"_loadingState"`
			Paidusage     bool   `json:"paidUsage"`
			Apipath       string `json:"apiPath"`
			Kalturaconfig struct {
				Partnerid                   interface{} `json:"partnerId"`
				Webplayerplaykitid          interface{} `json:"webPlayerPlaykitId"`
				Webplayeranonymousplaykitid interface{} `json:"webPlayerAnonymousPlaykitId"`
			} `json:"kalturaConfig"`
		} `json:"platformSettings"`
		Progress struct {
		} `json:"progress"`
	} `json:"videoPlayer"`
	Appstate struct {
		Colormode string `json:"colorMode"`
		Contents  struct {
			UrnOrmBook9781492092384ChapterCh04HTML struct {
				Apiurl         string      `json:"apiUrl"`
				Contentformat  string      `json:"contentFormat"`
				Description    string      `json:"description"`
				Duration       float64     `json:"duration"`
				Title          string      `json:"title"`
				Originalformat interface{} `json:"originalFormat"`
			} `json:"urn:orm:book:9781492092384:chapter:ch04.html"`
		} `json:"contents"`
		Fontsize        int    `json:"fontSize"`
		Hosttype        string `json:"hostType"`
		Readerwidth     int    `json:"readerWidth"`
		Tableofcontents map[string]struct {
			Sections []struct {
				Contentformat  string      `json:"contentFormat"`
				Contentid      string      `json:"contentId"`
				Duration       int         `json:"duration"`
				Title          string      `json:"title"`
				Depth          int         `json:"depth"`
				Apiurl         string      `json:"apiUrl"`
				Originalformat interface{} `json:"originalFormat"`
				Ourn           string      `json:"ourn"`
				Parentid       string      `json:"parentId"`
				Previousitem   interface{} `json:"previousItem"`
				Nextitem       struct {
					Contentid string `json:"contentId"`
					Ourn      string `json:"ourn"`
					Title     string `json:"title"`
				} `json:"nextItem"`
				Isephemeral                bool   `json:"isEphemeral"`
				Hasnonephemeraldescendants bool   `json:"hasNonEphemeralDescendants"`
				Fragment                   string `json:"fragment"`
				Pages                      int    `json:"pages"`
			} `json:"sections"`
		} `json:"tableOfContents"`
		Titles map[string]struct {
			Title        string `json:"title"`
			Slug         string `json:"slug"`
			Apiurl       string `json:"apiUrl"`
			Ourn         string `json:"ourn"`
			Description  string `json:"description"`
			Pagecount    int    `json:"pageCount"`
			Duration     int    `json:"duration"`
			Contributors struct {
				Authors  []string      `json:"authors"`
				Curators []interface{} `json:"curators"`
			} `json:"contributors"`
			Originalformat  interface{} `json:"originalFormat"`
			Publicationdate string      `json:"publicationDate"`
			Publishers      []struct {
				Name        string `json:"name"`
				UUID        string `json:"uuid"`
				Description struct {
					Textplain string `json:"textPlain"`
					Texthtml  string `json:"textHtml"`
				} `json:"description"`
				Academicexcluded bool `json:"academicExcluded"`
			} `json:"publishers"`
			Resources []struct {
				URL         string `json:"url"`
				Description string `json:"description"`
			} `json:"resources"`
			Topics []struct {
				Name string `json:"name"`
				ID   string `json:"id"`
				Slug string `json:"slug"`
			} `json:"topics"`
			Certifications struct {
				Identifier   interface{} `json:"identifier"`
				Practiceexam interface{} `json:"practiceExam"`
				Guides       interface{} `json:"guides"`
			} `json:"certifications"`
		} `json:"titles"`
	} `json:"appState"`
}
