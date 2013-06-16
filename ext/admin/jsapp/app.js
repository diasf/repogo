/*
 * This file is provided for custom JavaScript logic that your HTML files might need.
 * Maqetta includes this JavaScript file by default within HTML pages authored in Maqetta.
 */
require([
	"dijit/registry",
	"dijit/layout/BorderContainer",
	"dijit/layout/TabContainer",
	"dijit/layout/ContentPane",
	"dojox/grid/DataGrid",
	"dojo/store/JsonRest",
	"dojo/data/ObjectStore",
	"dojo/i18n!./nls/messages.js",
	"dojo/domReady!"
], function (registry, BorderContainer, TabContainer, ContentPane, DataGrid, JsonRest, ObjectStore, messages) {
	"use strict";
	// create the BorderContainer and attach it to our appLayout div
	var appLayout = new BorderContainer({
			design: "headline"
		}, "appLayout"),
		// create the TabContainer
		centerPanel = new ContentPane({
			region: "center",
			id: "centerPanel",
			"class": "centerPanel"
		}),
		userStore = new ObjectStore({objectStore: new JsonRest({ target: "/users/" })});
 
	// add the TabContainer as a child of the BorderContainer
	appLayout.addChild(centerPanel);
	
	// create and add the BorderContainer edge regions
	appLayout.addChild(new ContentPane({
		region: "top",
		"class": "edgePanel",
		content: "Header content (top)"
    }));
				
	appLayout.addChild(new ContentPane({
		region: "left",
		id: "leftCol",
		"class": "edgePanel",
		content: "Sidebar content (left)",
		splitter: true
	}));
					
	centerPanel.addChild(new DataGrid({
		id: 'usersGrid',
		store: userStore,
		structure: [[
			{'name': messages.usergrid_id,       'field': 'id',       'width': '33%'},
			{'name': messages.usergrid_login,    'field': 'login',    'width': '33%'},
			{'name': messages.usergrid_password, 'field': 'password', 'width': '33%'}
		]]
    }));
	// start up and do layout
	appLayout.startup();
});
