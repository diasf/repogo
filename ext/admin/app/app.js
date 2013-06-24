/*global $*/
// repogo Admin initialization
// --------------------------
var rAdm = rAdm || {};

rAdm.service = angular.module("rAdm.service", []);
rAdm.service.factory("userService", ["$http", function ($http) {
	"use strict";
	return {
		load: function (callback) {
			// load this from server..
			$http.get('/users').success(function (data) {
				callback(data);
			});
		}
	};
}]);

rAdm.service.factory("menuService", [function () {
	"use strict";
	return {
		menu: {
			entries: [
				{ text: 'Security', header: true },
				{ text: 'Users', link: '#Users', selected: true },
				{ text: 'Other', link: '#Other' },
				{ divider: true },
				{ text: 'Help', link: '#Help' }
			]
		},
		getMenu: function () {
			return this.menu;
		}
	};
}]);

rAdm.directive = angular.module("rAdm.directive", []);
rAdm.directive.directive('rAdmMenu', [function () {
	"use strict";
	return function (scope, element, attrs) {
		var ul = $(element);
		scope.$watch(attrs.rAdmMenu, function (menu) {
			$.each(menu.entries, function (idx, entry) {
				var li = $('<li>');
				// class
				if (entry.selected) {
					li.addClass("active");
				} else if (entry.divider) {
					li.addClass("divider");
				} else if (entry.header) {
					li.addClass("nav-header");
				}
				// is link
				if (entry.link && entry.text) {
					li.append($('<a>').attr('href', entry.link).html(entry.text));
				} else if (entry.text) {
					li.html(entry.text);
				}
				ul.append(li);
			});
		});
	};
}]);

rAdm.filter = angular.module("rAdm.filter", []);

rAdm.main = angular.module("rAdm", ["rAdm.service", "rAdm.directive", "rAdm.filter"]);
rAdm.main.run(["userService", function (userService) {
	"use strict";
	// initialize the module here...
}]);
rAdm.main.controller("MenuCtrl", ["$scope", "menuService", function ($scope, menuService) {
	"use strict";
	$scope.menu = menuService.getMenu();
}]);
rAdm.main.controller("UserCtrl", ["$scope", "userService", function ($scope, userService) {
	"use strict";
	userService.load(function (data) {
		$scope.users = data;
	});
}]);
