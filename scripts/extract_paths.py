#!/usr/bin/env python3

import argparse
import sys
import yaml


def main():
    parser = argparse.ArgumentParser(
        description=(
            "Generate a `yq` instructions script to add an `operationId` "
            "field to all endpoints in a swagger file"
        ),
    )
    parser.add_argument(
        "swagger_file",
        type=argparse.FileType("r"),
        default=sys.stdin,
        nargs="?",
        help="Path to the input Swagger YAML file",
    )
    parser.add_argument(
        "instructions_file",
        type=argparse.FileType("w"),
        default=sys.stdout,
        nargs="?",
        help="Path to the output `yq` instructions YAML file",
    )

    args = vars(parser.parse_args())

    swagger_file = args["swagger_file"]
    instructions_file = args["instructions_file"]
    paths = yaml.safe_load(swagger_file).get("paths")
    instructions = []
    for path, endpoints in paths.items():
        for http_verb in endpoints:
            instructions.append(
                {
                    "command": "update",
                    "path": "paths.%s.%s.operationId" % (path, http_verb),
                    "value": "TODO",
                }
            )

    yaml.safe_dump(instructions, instructions_file)

if __name__ == "__main__":
    main()
