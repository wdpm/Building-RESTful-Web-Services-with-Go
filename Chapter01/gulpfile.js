var gulp = require("gulp");
var shell = require('gulp-shell');

// This compiles new binary with source change
gulp.task("install-binary", shell.task([
  'go install github.com/narenaryan/romanserver'
]));

// Second argument tells install-binary is a dependency for restart-supervisor
gulp.task("restart-supervisor", ["install-binary"],  shell.task([
  'supervisorctl restart romanserver'
]))

gulp.task('watch', function() {
  // Watch the source code for all changes
  gulp.watch("*", ['install-binary', 'restart-supervisor']);

});

gulp.task('default', ['watch']);

//npm install gulp gulp-shell
