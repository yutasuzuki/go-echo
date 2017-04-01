'use strict';
const gulp = require('gulp');
const postcss = require('gulp-postcss');
const imported = require("postcss-easy-import");
const nested = require("postcss-nested");
const autoprefixer = require('autoprefixer');
const stylelint = require('stylelint');
const cssnano = require('cssnano');
const fileInclude = require('gulp-file-include');

gulp.task('style',  () => {
  return gulp.src('src/css/*.css')
    .pipe(postcss([
      autoprefixer(),
      imported(),
      nested(),
      stylelint(),
      cssnano()
    ]))
    .pipe( gulp.dest('dist/css/') );
});

gulp.task('include', function() {
  gulp.src('src/html/*.html')
    .pipe(fileInclude())
    .pipe(gulp.dest('../template/'));
});