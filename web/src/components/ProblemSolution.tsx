import styles from './ProblemSolution.module.css';

export default function ProblemSolution() {
    return (
        <section className={styles.section}>
            <div className={styles.container}>
                <div className={styles.header}>
                    <h2>Coding is fun. Writing docs isn't.</h2>
                    <p>Every hour you spend writing documentation is an hour you aren't building features. Worst of all, manually written docs are outdated the moment you push code.</p>
                </div>

                <div className={styles.comparison}>
                    <div className={styles.problem}>
                        <h3>The Old Way</h3>
                        <ul className={styles.list}>
                            <li className={styles.bad}>❌ Tedious manual writing</li>
                            <li className={styles.bad}>❌ Instantly outdated</li>
                            <li className={styles.bad}>❌ Inconsistent formats</li>
                            <li className={styles.bad}>❌ Developers hate it</li>
                        </ul>
                    </div>
                    <div className={styles.solution}>
                        <div className={styles.badge}>DevDocs Way</div>
                        <h3>The DevDocs Way</h3>
                        <ul className={styles.list}>
                            <li className={styles.good}>✅ Zero manual effort</li>
                            <li className={styles.good}>✅ Auto-sync with code</li>
                            <li className={styles.good}>✅ Beautiful, standard output</li>
                            <li className={styles.good}>✅ Developers love it</li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>
    );
}
