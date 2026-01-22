pipeline {
    agent any

    environment {
        DOCKER_USER = 'sudhanshud100'
        IMAGE_NAME  = 'my-go-app'
        DOCKER_HUB_CREDS = 'docker-hub-creds' 
    }

    stages {
        stage('Checkout') {
            steps {
                // Ye stage GitHub se code download karegi
                checkout scm
            }
        }
        stage('Gitleaks code check'){
            steps{
                script{
                    echo "Scannig the code .."
                    sh "docker run --rm -v ${WORKSPACE}:/path zricethezav/gitleaks:latest detect \
                    --source=/path --no-git --report-format json --report-path /path/gitleaks-report.json --verbose --redact"   
                }
            }
        }
        stage('SonarQube Analysis') {
            steps {
                script {
                    // Ye tool automatic download hoga
                    def scannerHome = tool 'sonar-scanner'
                    
                    // Ye name 'SonarQube' wahi hai jo Step 3-C mein diya tha
                    withSonarQubeEnv('SonarQube') {
                        sh "${scannerHome}/bin/sonar-scanner \
                        -Dsonar.projectKey=my-go-app \
                        -Dsonar.sources=. \
                        -Dsonar.host.url=http://sonarqube-server:9000"
                    }
                }
            }
        }
        stage('Docker Build') {
            steps {
                script {
                    echo "Building Docker Image..."
                    // Build number ko version ki tarah use karenge (v1, v2...)
                    sh "docker build -t ${DOCKER_USER}/${IMAGE_NAME}:v${env.BUILD_ID} ."
                    sh "docker tag ${DOCKER_USER}/${IMAGE_NAME}:v${env.BUILD_ID} ${DOCKER_USER}/${IMAGE_NAME}:latest"
                }
            }
        }

        stage('Docker Hub Push') {
            steps {
                script {
                    // Jenkins credentials manager se login karega
                    withCredentials([usernamePassword(credentialsId: "${DOCKER_HUB_CREDS}", passwordVariable: 'DOCKER_PASS', usernameVariable: 'DOCKER_USER_ENV')]) {
                        sh "echo \$DOCKER_PASS | docker login -u \$DOCKER_USER_ENV --password-stdin"
                        sh "docker push ${DOCKER_USER}/${IMAGE_NAME}:v${env.BUILD_ID}"
                        sh "docker push ${DOCKER_USER}/${IMAGE_NAME}:latest"
                    }
                }
            }
        }

        stage('Cleanup') {
            steps {
                echo "Removing local images to save space..."
                sh "docker rmi -f ${DOCKER_USER}/${IMAGE_NAME}:v${env.BUILD_ID} || true"
                sh "docker rmi -f ${DOCKER_USER}/${IMAGE_NAME}:latest || true"
            }
        }
        stage('Full Stack Deploy') {
            steps {
                script {
                    echo "Deploying Full Stack App..."
                    // docker-compose command ab chalegi
                    sh "docker-compose down" // Purana saaf karo
                    sh "docker-compose up -d" // Naya up karo
                    sh "docker ps" // Verify karne ke liye
                }
            }
        }
    }
}