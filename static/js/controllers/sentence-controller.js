(function() {
    angular.module('IYP').controller('SentenceController', ['$routeParams', 'Sentence', function($routeParams, Sentence) {
        var controller = this;
        alert($routeParams.level);
        Sentence.getById($routeParams.level).success(function(data) {
           controller.sentence = data;
           controller.currentLevel = parseInt($routeParams.level);
           controller.nextLevel = parseInt($routeParams.level) + 1;
           controller.previousLevel = parseInt($routeParams.level) - 1;
            if (controller.nextLevel === 11 || controller.nextLevel == 21 || controller.nextLevel == 31) {
                controller.nextLevel = 'congrats';
            }
            if (controller.previousLevel == 0) {
                controller.previousLevel = 10;
            } else if (controller.previousLevel == 10) {
                controller.previousLevel = 20;
            } else if (controller.previousLevel == 20) {
                controller.previousLevel = 30;
            }
        });
    }]);
})();