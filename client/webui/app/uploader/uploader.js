angular.module('repApp.uploader', [])
    .config(['$stateProvider', '$urlRouterProvider', function ($stateProvider, $urlRouterProvider) {
        'use strict';
        $stateProvider.state('uploader', {
            parent: 'main.default',
            url: '/uploader',
            views: {
                "main@main": {
                    templateUrl: 'uploader/uploader.html',
                    controller: 'UploaderCtrl'
                }
            }
        });
    }])
    .controller('UploaderCtrl', ['$scope', function ($scope) {
        'use strict';
    }]);
