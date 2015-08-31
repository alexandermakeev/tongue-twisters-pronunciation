(function() {
    angular.module('IYP').controller('SentenceController', ['$routeParams', 'Sentence', function($routeParams, Sentence) {
        var controller = this;
        Sentence.getById($routeParams.level).success(function(data) {
           controller.sentence = data;
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