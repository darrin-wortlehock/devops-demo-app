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

def version
  sh('scmversion current')
  IO.read(File.join(File.dirname(__FILE__), 'VERSION')) rescue 'unknown'
end

task :prepare do
  mkdir 'build'
  mkdir 'dist'
end

desc 'build the golang app'
task :build => [:clean, :prepare, :test, :bump] do
  sh("go build -ldflags '-X main.Build=#{version}' -o build/devops-demo-app")
end

desc 'test the golang app'
task :test do
  sh('go test')
end

task :bump do
  sh('scmversion bump auto')
end

desc 'create deb package'
task :package => [:clean, :prepare, :test, :bump, :build]