#!/bin/bash

ENV_FILE=.env.local

while read LINE; do
    export $LINE 
done < $ENV_FILE
