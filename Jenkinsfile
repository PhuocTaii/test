pipeline {
    agent any

    stages {
        stage('Build Docker Image') {
            steps {
                withCredentials([file(credentialsId: 'PROD_ENV_JSON_FILE', variable: 'PROD_ENV_FILE')]) {
                    script {
                        sh 'docker rm -f ia_03_container || echo "No existing container to remove"'
                        sh 'docker rmi -f ia_03:latest || echo "No existing image to remove"'
                        
                        // Ensure the 'env' directory exists
                        sh 'mkdir -p env'
                        // Copy the secret file to the 'env' directory
                        sh 'cp $PROD_ENV_FILE env/env.production.json'
                        sh 'ls -a'
                        sh 'ls -a env'
                        // Build the Docker image with no cache
                        sh 'docker build --no-cache -t ia_03:latest .'
                        // Verify the image is built
                        sh 'docker images | grep ia_03'

                        sh '''
                            docker run \
                                --name ia_03_container \
                                --restart always \
                                -dp 8000:8000 \
                                ia_03:latest
                        '''

                        sh 'docker logs -f ia_03_container'
                    }
                }
            }
        }
    }
}
