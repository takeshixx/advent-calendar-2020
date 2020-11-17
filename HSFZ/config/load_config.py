import toml


def load_config(path):
    return toml.loads(open(path, "r").read())
