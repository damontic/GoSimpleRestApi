pipeline {
    agent any

    stages {
        stage('Get') {
            steps {
                echo 'Getting..'
                sh go get -u github.com/damontic/GoSimpleRestApi
            }
        }
        stage('Check Style') {
            steps {
                echo 'Checking Style ...'
                sh golint github.com/damontic/GoSimpleRestApid
            }
        }
        stage('Check Suspicious Code') {
            steps {
                echo 'Checking Suspicious Code..'
                sh go vet github.com/damontic/GoSimpleRestApid
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                ansiblePlaybook {
                    playbook install_go_simple_rest_api.yml
                    inventory inventory_file
                }
            }
        }
    }
}
