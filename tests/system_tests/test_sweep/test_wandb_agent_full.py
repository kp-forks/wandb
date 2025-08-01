"""Agent tests."""

import contextlib
import io
import os
import unittest.mock

import wandb
from wandb.apis.public import Api


def test_agent_basic(user):
    sweep_ids = []
    sweep_configs = []
    sweep_resumed = []

    sweep_config = {
        "name": "My Sweep",
        "method": "grid",
        "parameters": {"a": {"values": [1, 2, 3]}},
    }

    def train():
        run = wandb.init()
        sweep_ids.append(run.sweep_id)
        sweep_configs.append(dict(run.config))
        sweep_resumed.append(run.resumed)
        run.finish()

    sweep_id = wandb.sweep(sweep_config)

    wandb.agent(sweep_id, function=train, count=1)

    assert len(sweep_ids) == len(sweep_configs) == 1
    assert sweep_ids[0] == sweep_id
    assert sweep_configs[0] == {"a": 1}
    assert sweep_resumed[0] is False


def test_agent_config_merge(user):
    sweep_configs = []

    def train():
        run = wandb.init(config={"extra": 2})
        sweep_configs.append(dict(run.config))
        run.finish()

    sweep_config = {
        "name": "My Sweep",
        "method": "grid",
        "parameters": {"a": {"values": [1, 2, 3]}},
    }

    with unittest.mock.patch.dict(os.environ, {"WANDB_CONSOLE": "off"}):
        sweep_id = wandb.sweep(sweep_config)
        wandb.agent(sweep_id, function=train, count=1)

    assert len(sweep_configs) == 1
    assert sweep_configs[0] == {"a": 1, "extra": 2}


def test_agent_config_ignore(user):
    sweep_configs = []

    def train():
        run = wandb.init(config={"a": "ignored", "extra": 2})
        sweep_configs.append(dict(run.config))
        run.finish()

    sweep_config = {
        "name": "My Sweep",
        "method": "grid",
        "parameters": {"a": {"values": [1, 2, 3]}},
    }

    sweep_id = wandb.sweep(sweep_config)
    wandb.agent(sweep_id, function=train, count=1)

    assert len(sweep_configs) == 1
    assert sweep_configs[0] == {"a": 1, "extra": 2}


def test_agent_ignore_project_entity_run_id(user):
    sweep_entities = []
    sweep_projects = []
    sweep_run_ids = []

    project_name = "actual"
    public_api = Api()
    public_api.create_project(project_name, user)

    def train():
        run = wandb.init(entity="ign", project="ignored", id="also_ignored")
        sweep_projects.append(run.project)
        sweep_entities.append(run.entity)
        sweep_run_ids.append(run.id)
        run.finish()

    sweep_config = {
        "name": "My Sweep",
        "method": "grid",
        "parameters": {"a": {"values": [1, 2, 3]}},
    }
    sweep_id = wandb.sweep(sweep_config, project=project_name)
    wandb.agent(sweep_id, function=train, count=1, project=project_name)

    assert len(sweep_projects) == len(sweep_entities) == 1
    assert sweep_projects[0] == "actual"
    assert sweep_entities[0] == user
    assert sweep_run_ids[0] != "also_ignored"


def test_agent_exception(user):
    sweep_config = {
        "name": "My Sweep",
        "method": "grid",
        "parameters": {"a": {"values": [1, 2, 3]}},
    }

    def train():
        wandb.init()
        raise Exception("Unexpected error")

    sweep_id = wandb.sweep(sweep_config)

    captured_stderr = io.StringIO()
    with contextlib.redirect_stderr(captured_stderr):
        wandb.agent(sweep_id, function=train, count=1)

    stderr_lines = captured_stderr.getvalue().splitlines()

    # Traceback with exception should appear before we finish the run.
    patterns = ["Traceback", "Exception: Unexpected error", "wandb: Find logs at:"]
    current_pattern = 0

    for line in stderr_lines:
        if line.startswith(patterns[current_pattern]):
            current_pattern += 1
            if current_pattern == len(patterns):
                break

    # Verify all patterns were found in order
    assert current_pattern == len(patterns), (
        f"Not found in stderr: '{patterns[current_pattern]}'"
    )
