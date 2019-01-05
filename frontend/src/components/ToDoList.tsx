import * as React from 'react';
import axios from 'axios';
import { createStyles, withStyles, WithStyles, Theme } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ToDoListItem from './ToDoListItem';
import * as config from '../config.json';

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
}
interface ToDoListProps extends WithStyles<typeof styles> { };
interface ToDoListState {
    toDos: ToDo[];
    activeId: string;
}

class ToDoList extends React.Component<ToDoListProps, ToDoListState> {
    state: ToDoListState = {
        toDos: [],
        activeId: '',
    };

    constructor(props: ToDoListProps) {
        super(props);

        this.handleToggleActive = this.handleToggleActive.bind(this);
        this.handleToggleComplete = this.handleToggleComplete.bind(this);
    }

    handleToggleComplete(toDoId: string): () => void {
        return () => {
            const newToDos: ToDo[] = [];
            let newActiveId = this.state.activeId;
            axios.post(`${config.apiEndpoint}/to-do/${toDoId}/toggle-complete`)
                .then((response) => {
                    console.log('response', response.data);
                    response.data.forEach((toDo: ToDo) => {
                        newToDos.push(toDo as ToDo);
                    });

                    this.setState({
                        toDos: newToDos,
                        activeId: newActiveId,
                    });
                });
        };
    }

    handleToggleActive(toDoId: string): () => void {
        return () => {
            const { activeId } = this.state;

            this.setState({
                activeId: (toDoId !== activeId) ? toDoId : '',
            });
        }
    }

    componentDidMount() {
        axios.get(`${config.apiEndpoint}/to-do`).then((response) => {
            const newToDos: ToDo[] = [];
            response.data.forEach((toDo: ToDo) => {
                newToDos.push(toDo as ToDo);
            });
            this.setState({ toDos: newToDos });
            console.log(response.data);
        });
    }

    render() {
        const { classes } = this.props;
        const { toDos, activeId } = this.state;

        return (
            <List className={classes.root}>
                {toDos.map((toDo) => (
                    <ToDoListItem
                        key={toDo.id}
                        toDo={toDo}
                        active={activeId === toDo.id}
                        handleToggleComplete={this.handleToggleComplete}
                        handleToggleActive={this.handleToggleActive}
                    />
                ))}
            </List>
        );
    }
}

export default withStyles(styles)(ToDoList);
