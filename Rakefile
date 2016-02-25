require 'fpm'
require 'fpm/rake_task'
require 'rake/clean'
CLEAN.include 'build', 'dist'

FPM::RakeTask.new(:package, { source: :dir, target: :deb, directory: '.'}) do |deb|
  deb.name = 'devops-demo-app'
  deb.package = 'dist/'
  deb.args = ['build/devops-demo-app=/usr/bin/devops-demo-app']
  deb.deb_upstart = 'upstart/devops-demo-app'
  deb.deb_no_default_config_files = false
end

task :prepare do
  mkdir 'build'
  mkdir 'dist'
end

task :build do
  puts 'build the golang code'
  sh('go build -o build/devops-demo-app')
end

desc 'create deb package'
task :package => [:clean, :prepare, :build]