var gulp	= require('gulp');
var sass 	= require('gulp-sass');
var bs 		= require('browser-sync').create();
var connect = require('gulp-connect');

gulp.task('browser-sync', function() {
	bs.init({
		server: {
			baseDir: './'
		},
		notify: false,
		port: '1550',
		injectChanges: true
	});
});

gulp.task('dev', ['browser-sync'], function() {
	gulp.watch('*.*').on('change', bs.reload);
	gulp.watch('app/**/*.*').on('change', bs.reload);
	gulp.watch('assets/**/*.*').on('change', bs.reload);
});

gulp.task('webserver', function() {
  	connect.server({
  		port: 1550
	});
});
 
gulp.task('default', ['webserver']);