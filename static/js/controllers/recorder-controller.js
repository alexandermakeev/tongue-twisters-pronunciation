(function() {
    angular.module("IYP").controller('RecorderController', function($scope, $compile) {
        this.timer = 3000;
        this.setTimer = function(timer) {
            this.timer = timer;
        };
        this.record = function(level) {
            var sound = new Audio("lib/record_start.wav");
            sound.play();
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

            recorder && recorder.record();

            setTimeout(function() {
                var sound = new Audio("lib/record_stop.wav");
                sound.play();
                recorder && recorder.stop();
                post(level);
                recorder && recorder.clear();
            }, this.timer);
        };
        function post(level) {
            recorder && recorder.exportWAV(function(blob) {
                blob.lastModifiedDate = new Date();
                blob.name = "file";

                var data = new FormData();
                data.append('file', blob);
                $.ajax({
                    url: baseUrl + '/api/translate/' + level,
                    data: data,
                    cache: false,
                    contentType: false,
                    processData: false,
                    type: 'POST',
                    success: function(data){
                        var div = document.getElementById("inner");
                        var response = document.createElement('div');
                        if (data.right == true) {
                            response.innerHTML = '<div class="bg-success">Yay! You are right!</div>';
                        } else {
                            response.innerHTML = '<div class="bg-danger">You said: <i>"' + data.sentence + '"</i></div>';
                        }
                        div.appendChild(response);
                        $compile(response)($scope);
                    }
                });
            });
        }
    })
})();