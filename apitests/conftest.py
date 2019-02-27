from os import path
from logging import config
import yaml


with open(path.join(path.dirname(path.realpath(__file__)), "log_spec.yaml"), "r") as log_spec_file:
    config.dictConfig(yaml.load(log_spec_file))