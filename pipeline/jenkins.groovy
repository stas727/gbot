pipeline {
    agent any
    parameters {
        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        choice(name: 'Arch', choices: ['amd64', 'arm64'], description: 'Pick Arch')
    }
    stages {
        stage('clone') {
            steps {
                echo 'cone repository'
                git branch: "${BRANCH}", url : "${REPO}"
            }
        }

        stage('test') {
            steps {
                echo "start testing"
                sh 'make test'
            }
        }

        stage('build') {
            steps {
                echo "build app"
                sh 'make build'
            }
        }
        stage('image') {
             steps {
                echo "build docker image"
                sh 'make image'
             }
        }

        stage('push') {
              steps {
                echo 'push to registry'
                script {
                    docker.withRegistry('', 'dockerHub') {
                          sh 'make push'
                      }
                }
              }
        }
    }
}