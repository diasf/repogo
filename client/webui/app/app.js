angular.
    module('repApp', [
        'ui.router',
        'ui.bootstrap',
        'repApp.main',
        'repApp.uploader'
    ]).
    config(['$urlRouterProvider', function ($urlRouterProvider) {
        'use strict';
        $urlRouterProvider.otherwise("/");
    }]);
