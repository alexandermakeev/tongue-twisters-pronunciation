(function() {
    angular.module('IYP').controller('PhraseController', ['$routeParams', 'Phrase', function($routeParams, Phrase) {
        var controller = this;
        Phrase.getById($routeParams.level).success(function(data) {
           controller.phrase = data;
           controller.currentLevel = parseInt($routeParams.level);
           controller.nextLevel = parseInt($routeParams.level) + 1;
           controller.previousLevel = parseInt($routeParams.level) - 1;
            if (controller.nextLevel > 15) {
                controller.nextLevel = 'congrats';
            }
            if (controller.previousLevel < 1) {
                controller.previousLevel = 15;
            }
        });
    }]);
})();