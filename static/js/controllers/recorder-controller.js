(function() {
    angular.module("IYP").controller('RecorderController', function($scope, $compile) {
        this.timer = 7000;
        this.setTimer = function(timer) {
            this.timer = timer;
        };
        this.init = function() {
            Recorder.record({
                start: function() {}
            });
            Recorder.stop();
        };
        this.record = function(level) {
            var time = this.timer / 1000;
            var i = time;
            var counterBack = setInterval(function(){
                i--;
                if(i>=0){
                    $('.progress-bar').css('width', (i/time)*100+'%');
                } else {
                    clearTimeout(counterBack);
                }

            }, 1000);

            Recorder.record({
                start: function() {}
            });
            setTimeout(function() {
                Recorder.stop();
                Recorder.upload({
                    url:        "http://localhost:9999/api/translate/" + level,
                    audioParam: "file",
                    success: function(data){
                        var parsed = JSON.parse(data);
                        var div = document.getElementById("inner");
                        var response = document.createElement('div');
                        if (parsed.right == true) {
                            response.innerHTML = '<div class="bg-success">Yay! You are right!</div>';
                        } else {
                            response.innerHTML = '<div class="bg-danger">You said: <i>"' + parsed.phrase + '"</i></div>';
                        }
                        div.appendChild(response);
                        $compile(response)($scope);
                    }
                });
            }, this.timer);

        };
    })
})();