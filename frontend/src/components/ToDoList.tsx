import * as React from 'react';
import axios from 'axios';
import { createStyles, withStyles, WithStyles, Theme } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ToDoListItem from './ToDoListItem';
import config from '../config';

const styles = (theme: Theme) => createStyles({
    root: {
        width: '100%',
        margin: 'auto',
    },
});

export interface ToDo {
    id: string;
    text: string;
    complete: boolean;
    active: boolean;
}

interface ToDoListProps extends WithStyles<typeof styles> {
}

interface ToDoListState {
    toDos: ToDo[];
}

class ToDoList extends React.Component<ToDoListProps, ToDoListState> {
    state: ToDoListState = {
        toDos: [],
    };

    constructor(props: ToDoListProps) {
        super(props);

        this.handleToggleActive = this.handleToggleActive.bind(this);
        this.handleToggleComplete = this.handleToggleComplete.bind(this);
    }

    handleToggleComplete(toDoId: string): () => void {
        return () => {
            const newToDos: ToDo[] = [];
            axios.post(`${config.apiEndpoint}/to-do/${toDoId}/toggle-complete`)
                .then((response) => {
                    response.data.forEach((toDo: ToDo) => {
                        newToDos.push(toDo as ToDo);
                    });

                    this.setState({
                        toDos: newToDos,
                    });
                });
        };
    }

    handleToggleActive(toDoId: string): () => void {
        return () => {
            const newToDos: ToDo[] = [];
            axios.post(`${config.apiEndpoint}/to-do/${toDoId}/toggle-active`)
                .then((response) => {
                    response.data.forEach((toDo: ToDo) => {
                        newToDos.push(toDo as ToDo);
                    });

                    this.setState({
                        toDos: newToDos,
                    });
                });
        }
    }

    componentDidMount() {
        axios.get(`${config.apiEndpoint}/to-do`)
            .then((response) => {
                const newToDos: ToDo[] = [];
                response.data.forEach((toDo: ToDo) => {
                    newToDos.push(toDo as ToDo);
                });
                this.setState({ toDos: newToDos });
            });
    }

    render() {
        const { classes } = this.props;
        const { toDos } = this.state;

        return (
            <List className={classes.root}>
                {toDos.map((toDo) => (
                    <ToDoListItem
                        key={toDo.id}
                        toDo={toDo}
                        handleToggleComplete={this.handleToggleComplete}
                        handleToggleActive={this.handleToggleActive}
                    />
                ))}
            </List>
        );
    }
}

export default withStyles(styles)(ToDoList);
