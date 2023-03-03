import argparse
from json import load
import os
import subprocess
from unicodedata import name

import json
import jinja2

if __name__ == "__main__":
    parser: argparse.ArgumentParser = argparse.ArgumentParser()    
    
    parser.add_argument("--template", help="Path for the template file")
    parser.add_argument("--values-file", help="Input values file path to render the template with")
    args: dict[str, any] = parser.parse_args().__dict__
    
    loader: jinja2.FileSystemLoader = jinja2.FileSystemLoader("../resources/")
    environment: jinja2.Environment = jinja2.Environment(loader=loader)

    with open(args["values_file"]) as input:
        data: dict[str, any] = json.load(input)


    template = environment.get_template(f'{args["template"]}.j2')
    rendered: str = template.render(data)
    
    with open(f'../bin/{args["template"]}', "w+") as output:
        output.write(rendered)

    subprocess.run("mkdir test", shell=True)
    
 