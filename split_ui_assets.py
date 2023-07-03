#!/usr/bin/env python3

import cv2


def generate_tiles(image_path, out_folder, step_width, step_height):
    img = cv2.imread(image_path)
    height = img.shape[0]
    width = img.shape[1]
    data = img.copy()

    x1 = 0
    y1 = 0
    rank = 52

    for y in range(0, height, step_height):
        for x in range(0, width, step_width):
            if (height - y) < step_height or (width - x) < step_width:
                break

            y1 = y + step_height
            x1 = x + step_width
            tiles = None
            rank += 1

            # check whether the patch width or height exceeds the image width or height
            if x1 >= width and y1 >= height:
                x1 = width - 1
                y1 = height - 1
                tiles = data[y:y + step_height, x:x + step_width]
            elif y1 >= height:  # when patch height exceeds the image height
                y1 = height - 1
                tiles = data[y:y + step_height, x:x + step_width]
            elif x1 >= width:  # when patch width exceeds the image width
                x1 = width - 1
                tiles = data[y:y + step_height, x:x + step_width]
            else:
                tiles = data[y:y + step_height, x:x + step_width]

            cv2.imwrite(f'{out_folder}/tile{rank}.png', tiles)
            cv2.rectangle(img, (x, y), (x1, y1), (0, 255, 0), 1)

# TODO: update the PNGs to make these work


print('Generating tiles for cards...')
generate_tiles('ui/src/assets/images/cards.png', 'ui/src/assets/images/cards', 0, 0)

print('Generating tiles for card runes...')
generate_tiles('ui/src/assets/images/cards-runes.png', 'ui/src/assets/images/card-runes', 0, 0)
