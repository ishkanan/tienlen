#!/usr/bin/env python3

import os
import shutil

import cv2


def rune_rank_gen():
    ranks = \
        [15, 14, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16] + \
        [2, 1, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3] + \
        [41, 40, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42] + \
        [28, 27, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29]
    for rank in ranks:
        yield rank
    raise StopIteration


def card_rank_gen():
    ranks = \
        ['-back-black', 28, 27, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29] + \
        ['-back-red', 15, 14, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16] + \
        [0, 2, 1, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3] + \
        [0, 41, 40, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42]
    for rank in ranks:
        yield rank
    raise StopIteration


def generate_tiles(image_path, out_folder, step_width, step_height, rank_gen):
    shutil.rmtree(out_folder, True)
    os.mkdir(out_folder)

    img = cv2.imread(image_path)
    height = img.shape[0]
    width = img.shape[1]
    data = img.copy()

    x1 = 0
    y1 = 0

    for y in range(0, height, step_height):
        for x in range(0, width, step_width):
            if (height - y) < step_height or (width - x) < step_width:
                break

            y1 = y + step_height
            x1 = x + step_width
            tiles = None
            rank = next(rank_gen)

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

            out_file = f'{out_folder}/tile{rank}.png'
            cv2.imwrite(out_file, tiles)
            cv2.rectangle(img, (x, y), (x1, y1), (0, 255, 0), 1)
            print(f'> created {out_file}')


print('Generating tiles for cards...')
generate_tiles('ui/src/assets/images/cards.png', 'ui/src/assets/images/cards', 80, 120, card_rank_gen())
os.remove('ui/src/assets/images/cards/tile0.png')

print('Generating tiles for card runes...')
generate_tiles('ui/src/assets/images/cards-runes.png', 'ui/src/assets/images/card-runes', 21, 32, rune_rank_gen())
