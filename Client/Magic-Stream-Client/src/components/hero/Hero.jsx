const Hero = () => {
  return (
    <div
      style={{
        height: "60vh",
        background:
          "linear-gradient(rgba(0,0,0,.7), rgba(0,0,0,.8)), url('https://images.unsplash.com/photo-1489599849927-2ee91cede3ba') center/cover",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        color: "white",
        textAlign: "center",
      }}
    >
      <div>
        <h1 style={{ fontSize: "4rem", fontWeight: "bold" }}>
          Magic Stream
        </h1>

        <p style={{ fontSize: "1.3rem" }}>
          Watch trailers. Discover movies. Get recommendations.
        </p>

        <button
          className="btn btn-danger btn-lg mt-3"
          onClick={() =>
            window.scrollTo({
              top: 700,
              behavior: "smooth",
            })
          }
        >
          Browse Movies
        </button>
      </div>
    </div>
  );
};

export default Hero;