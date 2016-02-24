require 'fpm'
require 'fpm/rake_task'

FPM::RakeTask.new(:package,
                  {
                      source: :dir,
                      target: :deb,
                      directory: File.join(File.expand_path(File.dirname(__FILE__)), 'dist')
                  })