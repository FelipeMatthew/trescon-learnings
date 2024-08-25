import cv2 as cv

img = cv.imread('images/animais02.jpg')

cv.imshow('Animais', img)

# close w/ any key
cv.waitKey(0)