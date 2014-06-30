/*
 * Localization Component
 */
angular.module('comp.localization', [])
    .factory('localize', ['$http', '$rootScope', '$window', '$filter', function ($http, $rootScope, $window, $filter) {
		"use strict";
		var language = '',
	        languagePending = '',
			dictionary = [],
			// flag to indicate if the service is currently loading a the resource
			resourceFileLoading = false,
			successCallback = function (data) {
                dictionary = data;
                resourceFileLoading = false;
				language = languagePending; // broadcast that a new resource file has been loaded
                $rootScope.$broadcast('localizeResourcesUpdated');
            },
			// loads an language resource file from the server
			initLocalizedResources = function (newLanguage) {
				if (!resourceFileLoading && language !== newLanguage) {
					resourceFileLoading = true;
					languagePending = newLanguage;
					// request the resource file
					$http({ method: "GET", url: 'messages/msg_' + languagePending + '.js', cache: false }).success(successCallback).error(function () {
						// request the default resource file
						languagePending = 'default';
						$http({ method: "GET", url: 'messages/msg_default.js', cache: false }).success(successCallback).error(function () {
							resourceFileLoading = false;
						});
					});
				}
			};
        // do the first load of the resource file
        initLocalizedResources($window.navigator.userLanguage || $window.navigator.language);
        return {
			// setting of language on the fly
			setLanguage: function (newLanguage) {
				if (language !== newLanguage) {
					initLocalizedResources(newLanguage);
				}
			},
			// localize a string 
			getLocalizedString: function (value, params) {
				// default the result to the string to be encoded
				var result = value,
					entry,
					idx;
				// make sure the dictionary has valid data
				if (dictionary.length > 0) {
					// use the filter service to only return those entries which match the value
					// and only take the first result
					entry = $filter('filter')(dictionary, function (element) { return element.key === value; })[0];
					if (entry && entry.value) {
						result = entry.value;
						// replace params if any
						if (params && params.length >= 1) {
							for (idx = 0; idx < params.length; idx++) {
								result = result.replace('{' + idx + '}', params[idx]);
							}
						}
					}
				}
				return result;
			}
		};
    }])
    // translation filter
    // {{ TOKEN | i18n }}
    .filter('i18n', ['localize', function (localize) {
		"use strict";
        return function (input) {
            return localize.getLocalizedString(input);
        };
    }])
    // translation directive : updates the text value of the attached element
    // <p data-i18n="TOKEN|PARAM1|PARAM2|.." >UPDATED TEXT</p>
    .directive('i18n', ['localize', function (localize) {
		"use strict";
        var i18nDirective = {
            restrict: "EAC",
            updateText: function (elm, token) {
                var values = token.split('|');
                if (values.length >= 1) {
                    elm.text(localize.getLocalizedString(values[0], values.slice(1)));
                }
            },
            link: function (scope, elm, attrs) {
                scope.$on('localizeResourcesUpdated', function () {
                    i18nDirective.updateText(elm, attrs.i18n);
                });
                attrs.$observe('i18n', function (value) {
                    i18nDirective.updateText(elm, attrs.i18n);
                });
            }
        };
        return i18nDirective;
    }]);
