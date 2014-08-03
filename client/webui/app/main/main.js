/**
 * The Main screen with menus and stuff...
 */
angular.module('repApp.main', [])
    .config(['$stateProvider', '$urlRouterProvider', function ($stateProvider, $urlRouterProvider) {
        'use strict';
        $stateProvider.state('main', {
            url: '/',
            abstract: true,
            templateUrl: 'main/main.html',
            controller: 'MainCtrl'
        });
        $stateProvider.state('main.default', {
            url: '',
            views: {
                "main": {
                    templateUrl: 'main/default.main.html'
                },
                "topMenu": {
                    templateUrl: 'main/default.topmenu.html'
                },
                "leftMenu": {
                    templateUrl: 'main/default.leftmenu.html'
                }
            }
        });
    }])
    .controller('MainCtrl', ['$scope', function ($scope) {
        'use strict';
        $scope.items = [
            'The first choice!',
            'And another choice for you.',
            'but wait! A third!'
        ];

        $scope.status = {
            isopen: false
        };

        $scope.toggled = function (open) {
        };

        $scope.toggleDropdown = function ($event) {
            $event.preventDefault();
            $event.stopPropagation();
            $scope.status.isopen = !$scope.status.isopen;
        };
    }]);
