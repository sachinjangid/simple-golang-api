pipeline {
    agent any
    
    environment {
        DOCKER_REGISTRY = 'docker.io'
        DOCKER_USERNAME = 'sachinjangid'
        PROJECT_NAME = 'simple-golang-api'
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Dependencies') {
            steps {
                sh 'go mod download'
            }
        }
        
        stage('Format Check') {
            steps {
                sh '''
                    if [ -n "$(gofmt -l .)" ]; then
                        echo "Code needs formatting!"
                        gofmt -l .
                        exit 1
                    fi
                    echo "Formatting OK!"
                '''
            }
        }
        
        stage('Tests') {
            steps {
                sh '''
                    go test -v -race ./...
                '''
            }
        }
        
        stage('Build') {
            steps {
                sh '''
                    go build -o myapp .
                '''
            }
        }
        
        stage('Docker Build') {
            steps {
                script {
                    // Define image tags
                    def imageName = "${env.DOCKER_USERNAME}/${env.PROJECT_NAME}"
                    def versionTag = "${imageName}:${env.BUILD_NUMBER}"
                    def latestTag = "${imageName}:latest"
                    
                    echo "Building Docker image: ${versionTag}"
                    
                    sh """
                    docker build -t ${versionTag} .
                    docker tag ${versionTag} ${latestTag}
                    """
                }
            }
        }
        
        stage('Docker Push') {
            steps {
                script {
                    def imageName = "${env.DOCKER_USERNAME}/${env.PROJECT_NAME}"
                    def versionTag = "${imageName}:${env.BUILD_NUMBER}"
                    def latestTag = "${imageName}:latest"
                    
                    withCredentials([usernamePassword(
                        credentialsId: 'docker-hub-credentials',  // Match the ID you set in Jenkins
                        usernameVariable: 'DOCKER_USER',
                        passwordVariable: 'DOCKER_PASS'
                    )]) {
                        sh """
                        # Login to Docker Hub
                        echo \$DOCKER_PASS | docker login -u \$DOCKER_USER --password-stdin
                        
                        # Push both tags
                        docker push ${versionTag}
                        docker push ${latestTag}
                        
                        # Logout (optional)
                        docker logout
                        """
                    }
                }
            }
        }
    }
    
    post {
        always {
            sh '''
                # Cleanup
                rm -f myapp coverage.out 2>/dev/null || true
                
                # Remove Docker images to save space
                docker images | grep "${DOCKER_USERNAME}/${PROJECT_NAME}" | awk '{print \$3}' | xargs -r docker rmi -f 2>/dev/null || true
            '''
        }
        success {
            echo "✅ Pipeline executed successfully!"
            script {
                def imageUrl = "https://hub.docker.com/r/${env.DOCKER_USERNAME}/${env.PROJECT_NAME}"
                echo "Docker image pushed: ${imageUrl}"
            }
        }
        failure {
            echo "❌ Pipeline failed!"
        }
    }
}