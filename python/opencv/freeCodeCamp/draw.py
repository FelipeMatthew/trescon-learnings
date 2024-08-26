import cv2 as cv
import numpy as np

blank = np.zeros((500,500,3), dtype='uint8')
cv.imshow('Blank', blank)

# 1. painting the entire image
# blank[200:300, 300:400] = 0,255,0
# cv.imshow('Green', blank)

# 2. Draw an rectangle
cv.rectangle(blank, (50,50), (250,250), (0,255,0), thickness=2)
cv.imshow('Retangle', blank)

# 3. Draw circle
cv.circle(blank,(blank.shape[1]//2, blank.shape[0]//2), 40,(0,0,255), thickness=3)
cv.imshow('Circle', blank)

# 4. Draw Line
cv.circle(blank, (0,0), (blank.shape[1]//2, blank.shape[0]//2), (0,0,255), thickness=3)
cv.imshow('Line', blank)

cv.waitKey(0)