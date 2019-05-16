# gtrchive

The **G**olang **T**witter a**R**chive is a simple Go program to ingest Tweets and store them in RethinkDB.

## Usage

The easist way to get up and running quickly is to use the official Docker [image][docker_image].
Be sure to include your Twitter API credentials, as well as the keywords to track.
See the official Twitter [documentation][twitter_docs] for the keyword format.

```bash
docker run -it --rm \
    -e GTR_TWITTER_ACCESS_TOKEN=<SECRET> \
    -e GTR_TWITTER_ACCESS_SECRET=<SECRET> \
    -e GTR_TWITTER_CONSUMER_KEY=<SECRET> \
    -e GTR_TWITTER_CONSUMER_SECRET=<SECRET> \
    -e GTR_TWITTER_TRACK='foo bar,baz' \
    jmckind/gtrchive:latest
```

## Development

Check out the source code locally and navigate to the directory.

```bash
git checout <REPO>/gtrchive.git
cd gtrchive
```

Next, ensure that the dependencies are present.

```bash
dep ensure
```

Set environment variables with Twitter credentials.

```bash
export GTR_TWITTER_ACCESS_TOKEN=<SECRET>
export GTR_TWITTER_ACCESS_SECRET=<SECRET>
export GTR_TWITTER_CONSUMER_KEY=<SECRET>
export GTR_TWITTER_CONSUMER_SECRET=<SECRET>
```

Set environment variable with keywords to track.

```bash
export GTR_TWITTER_TRACK='foo bar,baz'
```

Run the application.

```bash
go run gtrchive.go
```

## Release

Build the Docker image.

```bash
docker build <REPO>/gtrchive:latest .
```

Run the Docker image locally.

```bash
docker run -it --rm -e <ENV VARS>... <REPO>/gtrchive:latest
```

## License

gtrchive is released under the Apache 2.0 license. See the [LICENSE][license_file] file for details.

[license_file]:./LICENSE
[twitter_docs]: https://developer.twitter.com/en/docs/tweets/filter-realtime/api-reference/post-statuses-filter.html
[docker_image]: https://hub.docker.com/r/jmckind/gtrchive
