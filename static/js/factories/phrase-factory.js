(function() {
    angular.module('IYP').factory('Phrase', function PhraseFactory($http) {
       return {
           getById: function(id) {
               return $http({method: 'GET', url: baseUrl + '/api/phrases/' + id})
           }
       }
    });
})();