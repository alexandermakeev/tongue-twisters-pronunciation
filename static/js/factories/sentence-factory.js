(function() {
    angular.module('IYP').factory('Sentence', function SentenceFactory($http) {
       return {
           getById: function(id) {
               return $http({method: 'GET', url: baseUrl + '/api/sentences/' + id})
           }
       }
    });
})();