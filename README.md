# image-manager
API for uploading, resizing and downloading images.
Accepting images in png, jpeg.

Start: 
  - docker-compose up

Endpoints:
  - /upload/file(png/jpeg)
  - /download/$id/$quality(100, 75, 50, 25)
  
Packages:
  - broker - producer and consumer for message broker.
  - config - set up local variables from enviroment.
  - controllers - manage sizing image.
  - handlers - handler functions for endpoints, manage API responses.
  - helpers - connecting to RMQ, declare channel and queue.
  - middlewares - message broker middleware.
  - routes - endpoints.
  - storage - set up storage, manage saving and getting images,
  - utils - generate uniq id and resize image functions.
