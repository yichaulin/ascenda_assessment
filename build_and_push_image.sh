VERSION=$1

if [ $VERSION ]; then
    echo "Building $VERSION image..."
else
    echo "Not Given VERSION."
fi

docker build -t yichaulin/ascenda_assessment:$VERSION .
docker push yichaulin/ascenda_assessment:$VERSION

docker tag yichaulin/ascenda_assessment:$VERSION yichaulin/ascenda_assessment:latest
docker push yichaulin/ascenda_assessment:latest