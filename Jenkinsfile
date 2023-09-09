pipeline {
    agent any
    tools{
        go 'go1.20'
    }

    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
    }

    stages{
        stage('Build') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go mod download'
            }
        }
    }

    stages('Test') {
        steps {
            echo 'Running Test'
            sh 'go test ./...'
        }
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            
        }
    }  
}