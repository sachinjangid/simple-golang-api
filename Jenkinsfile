pipeline {
    agent any
    
    environment {
        GO_VERSION = '1.25'
        DOCKER_REGISTRY = 'your-docker-registry.com'
        PROJECT_NAME = 'simple-golang-api'
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Setup Go') {
            steps {
                script {
                    // Install specific Go version if needed
                    sh """
                    go version
                    """
                }
            }
        }
        
        stage('Code Format Check') {
            steps {
                script {
                    sh '''
                    # Check if code is properly formatted
                    if [ -n "$(gofmt -l .)" ]; then
                        echo "Code is not properly formatted. Files needing formatting:"
                        gofmt -l .
                        exit 1
                    fi
                    echo "Code formatting check passed!"
                    '''
                }
            }
        }
        
        stage('Run Tests') {
            steps {
                script {
                    sh '''
                    # Run tests with coverage
                    go test -v -race -coverprofile=coverage.out ./...
                    go tool cover -html=coverage.out -o coverage.html
                    '''
                }
            }
            post {
                always {
                    // Publish test results
                    publishHTML([
                        allowMissing: false,
                        alwaysLinkToLastBuild: false,
                        keepAll: true,
                        reportDir: '.',
                        reportFiles: 'coverage.html',
                        reportName: 'Test Coverage Report'
                    ])
                }
            }
        }
        
        stage('Build') {
            steps {
                script {
                    sh '''
                    # Build the application
                    go build -o myapp .
                    '''
                }
            }
        }
        
        stage('Docker Build & Push') {
            environment {
                DOCKER_IMAGE = "${DOCKER_REGISTRY}/${PROJECT_NAME}:${env.BUILD_NUMBER}"
                DOCKER_IMAGE_LATEST = "${DOCKER_REGISTRY}/${PROJECT_NAME}:latest"
            }
            steps {
                script {
                    sh """
                    # Build Docker image
                    docker build -t ${DOCKER_IMAGE} .
                    docker tag ${DOCKER_IMAGE} ${DOCKER_IMAGE_LATEST}
                    
                    # Push to registry (ensure credentials are set up in Jenkins)
                    docker push ${DOCKER_IMAGE}
                    docker push ${DOCKER_IMAGE_LATEST}
                    """
                }
            }
        }
    }
    
    post {
        always {
            // Cleanup
            sh 'rm -f myapp coverage.out coverage.html'
            
            // Notifications
            emailext (
                subject: "Build #${env.BUILD_NUMBER} - ${currentBuild.currentResult}",
                body: "Project: ${env.JOB_NAME}\nBuild: ${env.BUILD_NUMBER}\nStatus: ${currentBuild.currentResult}\nURL: ${env.BUILD_URL}",
                to: "team@yourcompany.com"
            )
        }
        success {
            echo "Pipeline executed successfully!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}