(function() {
    angular.module('IYP').config(function($routeProvider) {
       $routeProvider
           .when('/', {
               templateUrl: 'templates/home/index.html'
           })
           .when('/congrats', {
               templateUrl: 'templates/congrats/index.html'
           })
           .when('/:level', {
               templateUrl: 'templates/sentence/index.html',
               controller: 'SentenceController',
               controllerAs: 'sentenceCtrl'
           })
           .otherwise({redirectTo: '/'});
    });
})();