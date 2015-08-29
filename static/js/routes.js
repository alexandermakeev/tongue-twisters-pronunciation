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
               templateUrl: 'templates/phrase/index.html',
               controller: 'PhraseController',
               controllerAs: 'phraseCtrl'
           })
           .otherwise({redirectTo: '/'});
    });
})();