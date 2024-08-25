import cv2 as cv

img = cv.imread('images/pessoas01.jpg')
cv.imshow('Pessoas', img) # open a window

def rescaleFrames(frame, scale=0.75):
  # 0 is height an 1 is width
  # used for videos, images and live videos
  width = int(frame.shape[1] * scale)
  heigth = int(frame.shape[0] * scale)
  dimensions = (width,heigth)
  
  return cv.resize(frame, dimensions, interpolation=cv.INTER_AREA)

def changeResolution(width, height):
  capture.set(3, width)
  capture.set(4, height)

capture = cv.VideoCapture('videos/gato.mp4')

# resized image
resized_image = rescaleFrames(img, scale=.4)
image = cv.imshow('Imagem', resized_image)

while True:
  isTrue, frame = capture.read()
  
  frame_resized = rescaleFrames(frame, scale=.3)
  
  cv.imshow('Video', frame)
  cv.imshow('Video resized', frame_resized)
  if cv.waitKey(20) & 0xFF==ord('x'):
    break
  

capture.release()
cv.destroyAllWindows()