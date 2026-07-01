import Button from 'react-bootstrap/Button'
import { Link } from 'react-router-dom';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
import {faCirclePlay} from '@fortawesome/free-solid-svg-icons';
import "./Movie.css";
import { motion } from "framer-motion";

const Movie = ({movie,updateMovieReview}) => {
    return (
        <motion.div
    className="col-md-4 mb-4"
    key={movie._id}
    whileHover={{
        y: -10,
        scale: 1.03
    }}
    transition={{
        duration: 0.25
    }}
>
            <Link
                to={`/stream/${movie.youtube_id}`}
                style={{ textDecoration: 'none', color: 'inherit' }}
            >
            <div className="card h-100 shadow-sm movie-card">
                <div style={{position:"relative"}}>
                    <img src={movie.poster_path} alt={movie.title} 
                        className="card-img-top"
                        style={{
                            objectFit: "contain",
                            height: "250px",
                            width: "100%"
                        }}
                    />
                    <span className="play-icon-overlay">
                            <FontAwesomeIcon icon={faCirclePlay} />
                    </span>
                </div>
                <div className = "card-body d-flex flex-column">
                    <h5 className ="card-title">{movie.title}</h5>
                    <div className="movie-meta">
                        {movie.genre?.[0]?.genre_name}
                    </div>
                </div>
                {movie.ranking?.ranking_name && (
                    <span
                      className={`ranking-badge ${
                        movie.ranking?.ranking_name?.toLowerCase()
                      }`}
                    >
                      {movie.ranking?.ranking_name}
                    </span>
                )}
                  {updateMovieReview && (
                        <Button
                            variant="outline-info"
                            onClick={e => {
                                e.preventDefault();
                                updateMovieReview(movie.imdb_id);
                            }}
                            className="m-3"
                        >
                            Review
                        </Button>
                    )}
            </div>
            </Link>
        </motion.div>
    )
}
export default Movie;