import cv2 as cv

# Images
# img = cv.imread('images/animais02.jpg')
# cv.imshow('Animais', img)

# Videos 
capture = cv.VideoCapture('videos/cachorro.mp4')

while True:
  # Frame by frame
  isTrue, frame = capture.read()
  cv.imshow('Video', frame)
  
  if cv.waitKey(20) & 0xFF==ord('x'):
    break

capture.release()
cv.destroyAllWindows()