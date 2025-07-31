CREATE TABLE
  users (
    id CHAR(36) NOT NULL                                                                  ,
    -- id_sekolah CHAR(46)                                                                ,
    `name` VARCHAR(101) NOT NULL                                                          ,
    email VARCHAR(70) NOT NULL                                                            ,
    `role` ENUM("user", "guru_volunteer", "admin_sekolah") NOT NULL                       ,
    `password` VARCHAR(100) NOT NULL                                                      ,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP                                ,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP   ,
    PRIMARY KEY (id)                                                                      ,
    UNIQUE KEY (email)
    -- FOREIGN KEY (id_sekolah) REFERENCES sekolah (id)
  );

-- CREATE TABLE
--   sekolah (
--     id CHAR(36) NOT NULL                                                               ,
--     `nama` VARCHAR(101) NOT NULL                                                       ,
--     alamat TEXT NOT NULL                                                               ,
--     status_verifikasi BOOLEAN DEFAULT FALSE                                            ,
--     `password` VARCHAR(21) NOT NULL                                                    ,
--     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP                             ,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     PRIMARY KEY (id)                                                                   ,
--     UNIQUE KEY (email)
--   );
-- CREATE TABLE
--   testimoni (
--     id CHAR(36) NOT NULL                                                               ,
--     id_user CHAR(36)                                                                   ,
--     tipe CHAR(36)                                                                      ,
--     isi VARCHAR(101) NOT NULL                                                          ,
--     status_moderasi TEXT NOT NULL                                                      ,
--     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP                             ,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     PRIMARY KEY (id)                                                                   ,
--   );
-- CREATE TABLE
--   donasi (
--     id CHAR(36) NOT NULL                                                               ,
--     id_user CHAR(36)                                                                   ,
--     jenis ENUM(uang, barang) NOT NULL                                                  ,
--     jumlah INT UNSIGNED NOT NULL                                                       ,
--     `status` ENUM(pending, verified, `distributed`)                                    ,
--     progres TINYINT UNSIGNED NOT NULL DEFAULT 0                                        ,
--     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP                             ,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     PRIMARY KEY (id)                                                                   ,
--   );
-- CREATE TABLE
--   guru_volunteer (
--     id CHAR(36) NOT NULL                                                               ,
--     id_user CHAR(36)                                                                   ,
--     jenis ENUM(uang, barang) NOT NULL                                                  ,
--     jumlah INT UNSIGNED NOT NULL                                                       ,
--     `status` ENUM(pending, verified, `distributed`)                                    ,
--     progres TINYINT UNSIGNED NOT NULL DEFAULT 0                                        ,
--     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP                             ,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     PRIMARY KEY (id)                                                                   ,
--   );