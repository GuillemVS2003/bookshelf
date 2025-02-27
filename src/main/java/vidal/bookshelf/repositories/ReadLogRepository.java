package vidal.bookshelf.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import vidal.bookshelf.models.ReadLog;

public interface ReadLogRepository extends JpaRepository<ReadLog, Long> {

}
