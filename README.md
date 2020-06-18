# go-calc
This is an poc to demonstrate the go application deploynment in kubernetes cluster. Originally is forked from https://github.com/edcast/go-calc.git.

# to build the image
docker build -t {tagname} ./

# to push the image 
docker login {url} -u {uid} -p {pass}
docker push {imagename:tag}
You may find the latest images in docker hub {sumitdatta2015/go-calc}

# For deployment in kubernetes
kubectl create -f go-calc.yml
this will create pv,pvc,service and deployment in kubernetes cluster 
